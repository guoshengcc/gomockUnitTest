package main

// ShowData 显示数据
func ShowData(id string, foo Ifoo) string {
	res, err := foo.GetData(id)
	if err != nil {
		return err.Error()
	}
	return res
}
