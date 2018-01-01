package main

func (e GMOF_CaiPiao_Month) Process() error {
	saveData_GMOF_CaiPiao_Month(e.AccId, e.Title, e.Total, e.SPLY, e.LP, e.Url, e.Attachid, map[string]interface{}{
		"Id":       e.AccId,
		"Title":    e.Title,
		"Total":    e.Total,
		"SPLY":     e.SPLY,
		"LP":       e.LP,
		"Url":      e.Url,
		"Attachid": e.Attachid,
	})
	return nil
}

func (e GMOF_CaiPiao_List) Process() error {
	saveData_GMOF_CaiPiao_List(e.AccId, e.Title, e.Type, e.Url, map[string]interface{}{
		"Id":    e.AccId,
		"Title": e.Title,
		"Type":  e.Type,
		"Total": e.Url,
	})
	return nil
}
