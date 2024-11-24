package config

type CubeRecipe struct {
	Name             string
	Items            []string
	PurchaseRequired bool
	PurchaseItems    []string
}

var (
	Recipes = []CubeRecipe{

		// Perfects
		{
			Name:  "Perfect Amethyst",
			Items: []string{"FlawlessAmethyst", "FlawlessAmethyst", "FlawlessAmethyst"},
		},
		{
			Name:  "Perfect Diamond",
			Items: []string{"FlawlessDiamond", "FlawlessDiamond", "FlawlessDiamond"},
		},
		{
			Name:  "Perfect Emerald",
			Items: []string{"FlawlessEmerald", "FlawlessEmerald", "FlawlessEmerald"},
		},
		{
			Name:  "Perfect Ruby",
			Items: []string{"FlawlessRuby", "FlawlessRuby", "FlawlessRuby"},
		},
		{
			Name:  "Perfect Sapphire",
			Items: []string{"FlawlessSapphire", "FlawlessSapphire", "FlawlessSapphire"},
		},
		{
			Name:  "Perfect Topaz",
			Items: []string{"FlawlessTopaz", "FlawlessTopaz", "FlawlessTopaz"},
		},
		{
			Name:  "Perfect Skull",
			Items: []string{"FlawlessSkull", "FlawlessSkull", "FlawlessSkull"},
		},

		// Token
		{
			Name:  "Token of Absolution",
			Items: []string{"TwistedEssenceOfSuffering", "ChargedEssenceOfHatred", "BurningEssenceOfTerror", "FesteringEssenceOfDestruction"},
		},

		// Runes
		{
			Name:  "Upgrade El",
			Items: []string{"ElRune", "ElRune", "ElRune"},
		},
		{
			Name:  "Upgrade Eld",
			Items: []string{"EldRune", "EldRune", "EldRune"},
		},
		{
			Name:  "Upgrade Tir",
			Items: []string{"TirRune", "TirRune", "TirRune"},
		},
		{
			Name:  "Upgrade Nef",
			Items: []string{"NefRune", "NefRune", "NefRune"},
		},
		{
			Name:  "Upgrade Eth",
			Items: []string{"EthRune", "EthRune", "EthRune"},
		},
		{
			Name:  "Upgrade Ith",
			Items: []string{"IthRune", "IthRune", "IthRune"},
		},
		{
			Name:  "Upgrade Tal",
			Items: []string{"TalRune", "TalRune", "TalRune"},
		},
		{
			Name:  "Upgrade Ral",
			Items: []string{"RalRune", "RalRune", "RalRune"},
		},
		{
			Name:  "Upgrade Ort",
			Items: []string{"OrtRune", "OrtRune", "OrtRune"},
		},
		{
			Name:  "Upgrade Thul",
			Items: []string{"ThulRune", "ThulRune", "ThulRune", "ChippedTopaz"},
		},
		{
			Name:  "Upgrade Amn",
			Items: []string{"AmnRune", "AmnRune", "AmnRune", "ChippedAmethyst"},
		},
		{
			Name:  "Upgrade Sol",
			Items: []string{"SolRune", "SolRune", "SolRune", "ChippedSapphire"},
		},
		{
			Name:  "Upgrade Shael",
			Items: []string{"ShaelRune", "ShaelRune", "ShaelRune", "ChippedRuby"},
		},
		{
			Name:  "Upgrade Dol",
			Items: []string{"DolRune", "DolRune", "DolRune", "ChippedEmerald"},
		},
		{
			Name:  "Upgrade Hel",
			Items: []string{"HelRune", "HelRune", "HelRune", "ChippedDiamond"},
		},
		{
			Name:  "Upgrade Io",
			Items: []string{"IoRune", "IoRune", "IoRune", "FlawedTopaz"},
		},
		{
			Name:  "Upgrade Lum",
			Items: []string{"LumRune", "LumRune", "LumRune", "FlawedAmethyst"},
		},
		{
			Name:  "Upgrade Ko",
			Items: []string{"KoRune", "KoRune", "KoRune", "FlawedSapphire"},
		},
		{
			Name:  "Upgrade Fal",
			Items: []string{"FalRune", "FalRune", "FalRune", "FlawedRuby"},
		},
		{
			Name:  "Upgrade Lem",
			Items: []string{"LemRune", "LemRune", "LemRune", "FlawedEmerald"},
		},
		{
			Name:  "Upgrade Pul",
			Items: []string{"PulRune", "PulRune", "FlawedDiamond"},
		},
		{
			Name:  "Upgrade Um",
			Items: []string{"UmRune", "UmRune", "Topaz"},
		},
		{
			Name:  "Upgrade Mal",
			Items: []string{"MalRune", "MalRune", "Amethyst"},
		},
		{
			Name:  "Upgrade Ist",
			Items: []string{"IstRune", "IstRune", "Sapphire"},
		},
		{
			Name:  "Upgrade Gul",
			Items: []string{"GulRune", "GulRune", "Ruby"},
		},
		{
			Name:  "Upgrade Vex",
			Items: []string{"VexRune", "VexRune", "Emerald"},
		},
		{
			Name:  "Upgrade Ohm",
			Items: []string{"OhmRune", "OhmRune", "Diamond"},
		},
		{
			Name:  "Upgrade Lo",
			Items: []string{"LoRune", "LoRune", "FlawlessTopaz"},
		},
		{
			Name:  "Upgrade Sur",
			Items: []string{"SurRune", "SurRune", "FlawlessAmethyst"},
		},
		{
			Name:  "Upgrade Ber",
			Items: []string{"BerRune", "BerRune", "FlawlessSapphire"},
		},
		{
			Name:  "Upgrade Jah",
			Items: []string{"JahRune", "JahRune", "FlawlessRuby"},
		},
		{
			Name:  "Upgrade Cham",
			Items: []string{"ChamRune", "ChamRune", "FlawlessEmerald"},
		},

		// Crafting
		{
			Name:  "Reroll GrandCharms",
			Items: []string{"GrandCharm", "Perfect", "Perfect", "Perfect"}, // Special handling in hasItemsForRecipe
		},

		{
			Name:  "Reroll SmallCharms",
			Items: []string{"SmallCharm", "Perfect", "Perfect", "Perfect"}, // Special handling in hasItemsForRecipe
		},

		{
			Name:  "Reroll Monarchs",
			Items: []string{"Monarch", "Perfect", "Perfect", "Perfect"}, // Special handling in hasItemsForRecipe
		},

		// Caster Amulet
		{
			Name:             "Caster Amulet",
			Items:            []string{"RalRune", "PerfectAmethyst", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Amulet"},
		},

		// Caster Ring
		{
			Name:             "Caster Ring",
			Items:            []string{"AmnRune", "PerfectAmethyst", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Ring"},
		},

		// Blood Gloves
		{
			Name:             "Blood Gloves",
			Items:            []string{"NefRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"HeavyGloves", "SharkskinGloves", "VampireboneGloves"},
		},

		// Blood Boots
		{
			Name:             "Blood Boots",
			Items:            []string{"EthRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"LightPlatedBoots", "BattleBoots", "MirroredBoots"},
		},

		// Blood Belt
		{
			Name:             "Blood Belt",
			Items:            []string{"TalRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Belt", "MeshBelt", "MithrilCoil"},
		},

		// Blood Helm
		{
			Name:             "Blood Helm",
			Items:            []string{"RalRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Helm", "Casque", "Armet"},
		},

		// Blood Armor
		{
			Name:             "Blood Armor",
			Items:            []string{"ThulRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"PlateMail", "TemplarPlate", "HellforgePlate"},
		},

		// Blood Weapon
		{
			Name:             "Blood Weapon",
			Items:            []string{"OrtRune", "PerfectRuby", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Axe"},
		},

		// Safety Shield
		{
			Name:             "Safety Shield",
			Items:            []string{"EthRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"KiteShield", "DragonShield", "Monarch"},
		},

		// Safety Armor
		{
			Name:             "Safety Armor",
			Items:            []string{"NefRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"BreastPlate", "Curiass", "GreatHauberk"},
		},

		// Safety Boots
		{
			Name:             "Safety Boots",
			Items:            []string{"OrtRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Greaves", "WarBoots", "MyrmidonBoots"},
		},

		// Safety Gloves
		{
			Name:             "Safety Gloves",
			Items:            []string{"RalRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Gauntlets", "WarGauntlets", "OgreGauntlets"},
		},

		// Safety Belt
		{
			Name:             "Safety Belt",
			Items:            []string{"TalRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Sash", "DemonhideSash", "SpiderwebSash"},
		},

		// Safety Helm
		{
			Name:             "Safety Helm",
			Items:            []string{"IthRune", "PerfectEmerald", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"Crown", "GrandCrown", "Corona"},
		},

		// Hitpower Gloves
		{
			Name:             "Hitpower Gloves",
			Items:            []string{"OrtRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"ChainGloves", "HeavyBracers", "Vambraces"},
		},

		// Hitpower Boots
		{
			Name:             "Hitpower Boots",
			Items:            []string{"RalRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"ChainBoots", "MeshBoots", "Boneweave"},
		},

		// Hitpower Belt
		{
			Name:             "Hitpower Belt",
			Items:            []string{"TalRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"HeavyBelt", "BattleBelt", "TrollBelt"},
		},

		// Hitpower Helm
		{
			Name:             "Hitpower Helm",
			Items:            []string{"NefRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"FullHelm", "Basinet", "GiantConch"},
		},

		// Hitpower Armor
		{
			Name:             "Hitpower Armor",
			Items:            []string{"EthRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"FieldPlate", "Sharktooth", "KrakenShell"},
		},

		// Hitpower Shield
		{
			Name:             "Hitpower Shield",
			Items:            []string{"IthRune", "PerfectSapphire", "Jewel"},
			PurchaseRequired: true,
			PurchaseItems:    []string{"GothicShield", "AncientShield", "Ward"},
		},
	}
)

// List of available recipe names for configuration
var AvailableRecipes = []string{
	// List out all the recipe names from cube_recipes.go
	"Perfect Amethyst",
	"Perfect Diamond",
	"Perfect Emerald",
	"Perfect Ruby",
	"Perfect Sapphire",
	"Perfect Topaz",
	"Perfect Skull",
	"Token of Absolution",
	"Upgrade El",
	"Upgrade Eld",
	"Upgrade Tir",
	"Upgrade Nef",
	"Upgrade Eth",
	"Upgrade Ith",
	"Upgrade Tal",
	"Upgrade Ral",
	"Upgrade Ort",
	"Upgrade Thul",
	"Upgrade Amn",
	"Upgrade Sol",
	"Upgrade Shael",
	"Upgrade Dol",
	"Upgrade Hel",
	"Upgrade Io",
	"Upgrade Lum",
	"Upgrade Ko",
	"Upgrade Fal",
	"Upgrade Lem",
	"Upgrade Pul",
	"Upgrade Um",
	"Upgrade Mal",
	"Upgrade Ist",
	"Upgrade Gul",
	"Upgrade Vex",
	"Upgrade Ohm",
	"Upgrade Lo",
	"Upgrade Sur",
	"Upgrade Ber",
	"Upgrade Jah",
	"Upgrade Cham",
	"Reroll GrandCharms",
	"Reroll SmallCharms",
	"Reroll Monarchs",
	"Caster Amulet",
	"Caster Ring",
}
