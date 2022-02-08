package town

import (
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/game/data"
	"github.com/hectorgimenez/koolo/internal/helper"
	"github.com/hectorgimenez/koolo/internal/hid"
	"time"
)

type Manager struct {
	towns       map[data.Area]Town
	pf          helper.PathFinder
	dr          data.DataRepository
	shopManager ShopManager
	actionChan  chan<- action.Action
}

func NewTownManager(repository data.DataRepository, pf helper.PathFinder, shopManager ShopManager, actionChan chan<- action.Action) Manager {
	return Manager{
		towns: map[data.Area]Town{
			data.AreaHarrogath: A5{},
		},
		pf:          pf,
		dr:          repository,
		actionChan:  actionChan,
		shopManager: shopManager,
	}
}

func (tm Manager) BuyPotionsAndTPs(area data.Area) {
	t := tm.getTownByArea(area)
	tm.pf.InteractToNPC(t.RefillNPC())
	tm.openTradeMenu()
	tm.shopManager.buyPotsAndTPs()
}

func (tm Manager) Repair(area data.Area) {
	t := tm.getTownByArea(area)
	tm.pf.InteractToNPC(t.RepairNPC())
	tm.openTradeMenu()
	tm.actionChan <- action.NewAction(
		action.PriorityNormal,
		action.NewMouseDisplacement(int(float32(hid.GameAreaSizeX)/3.52), int(float32(hid.GameAreaSizeY)/1.37), time.Millisecond*850),
		action.NewMouseClick(hid.LeftButton, time.Millisecond*1300),
	)
}

func (tm Manager) ReviveMerc(area data.Area) {
	t := tm.getTownByArea(area)
	tm.pf.InteractToNPC(t.MercContractorNPC())
	tm.openTradeMenu()
}

func (tm Manager) WPTo(act int, area int) {
	for _, o := range tm.dr.GameData().Objects {
		if o.IsWaypoint() {
			tm.pf.InteractToObject(o)
			return
		}
	}
}

func (tm Manager) openTradeMenu() {
	d := tm.dr.GameData()
	if d.OpenMenus.NPCInteract {
		tm.actionChan <- action.NewAction(
			action.PriorityNormal,
			action.NewKeyPress("down", time.Millisecond*760),
			action.NewKeyPress("enter", time.Second),
		)
		time.Sleep(time.Second)
	}

}

func (tm Manager) getTownByArea(area data.Area) Town {
	return tm.towns[area]
}