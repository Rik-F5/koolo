package action

import (
	"slices"
	"strings"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/nip"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/context"
	"github.com/hectorgimenez/koolo/internal/utils"
)

func CubeRecipes() error {
	ctx := context.Get()
	ctx.SetLastAction("CubeRecipes")

	// If cubing is disabled from settings just return nil
	if !ctx.CharacterCfg.CubeRecipes.Enabled {
		ctx.Logger.Debug("Cube recipes are disabled, skipping")
		return nil
	}

	itemsInStash := ctx.Data.Inventory.ByLocation(item.LocationStash, item.LocationSharedStash)
	for _, recipe := range config.Recipes {
		// Check if the current recipe is Enabled
		if !slices.Contains(ctx.CharacterCfg.CubeRecipes.EnabledRecipes, recipe.Name) {
			// is this really needed ? making huge logs
			//		ctx.Logger.Debug("Cube recipe is not enabled, skipping", "recipe", recipe.Name)
			continue
		}

		ctx.Logger.Debug("Cube recipe is enabled, processing", "recipe", recipe.Name)

		continueProcessing := true
		for continueProcessing {
			if items, hasItems := hasItemsForRecipe(ctx, itemsInStash, recipe); hasItems {

				// TODO: Check if we have the items in our storage and if not, purchase them, else take the item from the storage
				if recipe.PurchaseRequired {
					err := GambleSingleItem(recipe.PurchaseItems, item.QualityMagic)
					if err != nil {
						ctx.Logger.Error("Error gambling item, skipping recipe", "error", err, "recipe", recipe.Name)
						break
					}

					purchasedItem := getPurchasedItem(ctx, recipe.PurchaseItems)
					if purchasedItem.Name == "" {
						ctx.Logger.Debug("Could not find purchased item. Skipping recipe", "recipe", recipe.Name)
						break
					}

					// Add the purchased item the list of items to cube
					items = append(items, purchasedItem)
				}

				// Add items to the cube and perform the transmutation
				err := CubeAddItems(items...)
				if err != nil {
					return err
				}
				if err = CubeTransmute(); err != nil {
					return err
				}

				// Get a list of items that are in our invetory
				itemsInInv := ctx.Data.Inventory.ByLocation(item.LocationInventory)

				stashingRequired := false
				stashingGrandCharm := false

				// Check if the items that are not in the protected invetory slots should be stashed
				for _, item := range itemsInInv {
					// If item is not in the protected slots, check if it should be stashed
					if ctx.CharacterCfg.Inventory.InventoryLock[item.Position.Y][item.Position.X] == 1 {

						shouldStash, reason, _ := shouldStashIt(item, false)

						if shouldStash {
							ctx.Logger.Debug("Stashing item after cube recipe.", "item", item.Name, "recipe", recipe.Name, "reason", reason)
							stashingRequired = true
						} else if item.Name == "GrandCharm" || item.Name == "SmallCharm" || item.Name == "Monarch" {
							ctx.Logger.Debug("Checking if we need to stash a GrandCharm that doesn't match any NIP rules.", "recipe", recipe.Name)
							// Check if we have a GrandCharm in stash that doesn't match any NIP rules
							hasUnmatchedGrandCharm := false
							for _, stashItem := range itemsInStash {
								if stashItem.Name == item.Name {
									if _, result := ctx.CharacterCfg.Runtime.Rules.EvaluateAll(stashItem); result != nip.RuleResultFullMatch {
										hasUnmatchedGrandCharm = true
										break
									}
								}
							}
							if !hasUnmatchedGrandCharm {

								ctx.Logger.Debug("GrandCharm doesn't match any NIP rules and we don't have any in stash to be used for this recipe. Stashing it.", "recipe", recipe.Name)
								stashingRequired = true
								stashingGrandCharm = true

							} else {
								DropInventoryItem(item)
								utils.Sleep(500)
							}
						} else {
							DropInventoryItem(item)
							utils.Sleep(500)
						}
					}
				}

				// Add items to the stash if needed
				if stashingRequired && !stashingGrandCharm {
					_ = Stash(false)
				} else if stashingGrandCharm {
					// Force stashing of the invetory
					_ = Stash(true)
				}

				// Remove or decrement the used items from itemsInStash
				itemsInStash = removeUsedItems(itemsInStash, items)
			} else {
				continueProcessing = false
			}
		}
	}

	return nil
}

func hasItemsForRecipe(ctx *context.Status, items []data.Item, recipe config.CubeRecipe) ([]data.Item, bool) {

	// Special handling for "Reroll" recipes
	if strings.HasPrefix(recipe.Name, "Reroll") {
		return hasItemsForReroll(ctx, items, recipe)
	}

	recipeItems := make(map[string]int)
	for _, item := range recipe.Items {
		recipeItems[item]++
	}

	itemsForRecipe := []data.Item{}

	// Iterate over the items in our stash to see if we have the items for the recipie.
	for _, item := range items {
		if count, ok := recipeItems[string(item.Name)]; ok {

			// Let's make sure we don't use an item we don't want to. Add more if needed (depending on the recipes we have)
			if item.Name == "Jewel" {
				if _, result := ctx.CharacterCfg.Runtime.Rules.EvaluateAll(item); result == nip.RuleResultFullMatch {
					continue
				}
			}

			itemsForRecipe = append(itemsForRecipe, item)

			// Check if we now have exactly the needed count before decrementing
			count -= 1
			if count == 0 {
				delete(recipeItems, string(item.Name))
				if len(recipeItems) == 0 {
					return itemsForRecipe, true
				}
			} else {
				recipeItems[string(item.Name)] = count
			}
		}
	}

	// We don't have all the items for the recipie.
	return nil, false
}

func hasItemsForReroll(ctx *context.Status, items []data.Item, recipe config.CubeRecipe) ([]data.Item, bool) {
	var Item data.Item
	perfectGems := make([]data.Item, 0, 3)

	for _, itm := range items {
		// ctx.Logger.Debug("func hasItemsForReroll for item", "item", itm.Name)
		if string(itm.Name) == recipe.Items[0] {
			if _, result := ctx.CharacterCfg.Runtime.Rules.EvaluateAll(itm); result != nip.RuleResultFullMatch && itm.Quality == item.QualityMagic {
				Item = itm
			}
		} else if isPerfectGem(itm) && len(perfectGems) < 3 {
			perfectGems = append(perfectGems, itm)
		}

		if Item.Name != "" && len(perfectGems) == 3 {
			return append([]data.Item{Item}, perfectGems...), true
		}
	}

	return nil, false
}

func isPerfectGem(item data.Item) bool {
	perfectGems := []string{"PerfectAmethyst", "PerfectDiamond", "PerfectEmerald", "PerfectRuby", "PerfectSapphire", "PerfectTopaz", "PerfectSkull"}

	// Get context to check enabled recipes
	ctx := context.Get()

	// If recipes are enabled, remove gems that are mentioned in enabled recipes
	if ctx.CharacterCfg.CubeRecipes.Enabled {
		for _, recipe := range config.Recipes {
			// Skip if recipe is not enabled
			if !slices.Contains(ctx.CharacterCfg.CubeRecipes.EnabledRecipes, recipe.Name) {
				continue
			}

			// Check if any perfect gems are mentioned in recipe items
			for _, recipeItem := range recipe.Items {
				beforeLen := len(perfectGems)
				perfectGems = slices.DeleteFunc(perfectGems, func(gem string) bool {
					return gem == recipeItem
				})
				if len(perfectGems) != beforeLen {
					ctx.Logger.Debug("Removed gem from perfect gems list", "recipe", recipe.Name, "gem", recipeItem, "remaining_gems", perfectGems)
				}
			}
		}
	}

	// Check if the item is in the remaining perfect gems list
	isGemInList := slices.Contains(perfectGems, string(item.Name))
	return isGemInList
}

func removeUsedItems(stash []data.Item, usedItems []data.Item) []data.Item {
	remainingItems := make([]data.Item, 0)
	usedItemMap := make(map[string]int)

	// Populate a map with the count of used items
	for _, item := range usedItems {
		usedItemMap[string(item.Name)] += 1 // Assuming 'ID' uniquely identifies items in 'usedItems'
	}

	// Filter the stash by excluding used items based on the count in the map
	for _, item := range stash {
		if count, exists := usedItemMap[string(item.Name)]; exists && count > 0 {
			usedItemMap[string(item.Name)] -= 1
		} else {
			remainingItems = append(remainingItems, item)
		}
	}

	return remainingItems
}

func getPurchasedItem(ctx *context.Status, purchaseItems []string) data.Item {
	itemsInInv := ctx.Data.Inventory.ByLocation(item.LocationInventory)
	for _, citem := range itemsInInv {
		for _, pi := range purchaseItems {
			if string(citem.Name) == pi && citem.Quality == item.QualityMagic {
				return citem
			}
		}
	}
	return data.Item{}
}
