package logic

import (
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/stores"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

func ToPageInfo(info *dm.PageInfo, defaultOrders ...def.OrderBy) *def.PageInfo {
	if info == nil {
		return nil
	}

	var orders = defaultOrders
	if infoOrders := info.GetOrders(); len(infoOrders) > 0 {
		orders = make([]def.OrderBy, 0, len(infoOrders))
		for _, infoOd := range infoOrders {
			if infoOd.GetFiled() != "" {
				orders = append(orders, def.OrderBy{infoOd.GetFiled(), infoOd.GetSort()})
			}
		}
	}

	return &def.PageInfo{
		Page:   info.GetPage(),
		Size:   info.GetSize(),
		Orders: orders,
	}
}

func ToPageInfoWithDefault(info *dm.PageInfo, defau *def.PageInfo) *def.PageInfo {
	if page := ToPageInfo(info); page == nil {
		return defau
	} else {
		if page.Page == 0 {
			page.Page = defau.Page
		}
		if page.Size == 0 {
			page.Size = defau.Size
		}
		if len(page.Orders) == 0 {
			page.Orders = defau.Orders
		}
		return page
	}
}

func ToDmPoint(point *stores.Point) *dm.Point {
	if point == nil {
		return nil
	}
	return &dm.Point{Longitude: point.Longitude, Latitude: point.Latitude}
}
func ToStorePoint(point *dm.Point) stores.Point {
	if point == nil {
		return stores.Point{Longitude: 0, Latitude: 0}
	}
	return stores.Point{Longitude: point.Longitude, Latitude: point.Latitude}
}
func ToDeviceCores(in []*dm.DeviceCore) []*devices.Core {
	list := make([]*devices.Core, 0, len(in))
	for _, v := range in {
		if v == nil {
			continue
		}
		list = append(list, &devices.Core{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
		})
	}
	return list
}
