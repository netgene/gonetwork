package business

type Ts_order_etr struct {
	Etrorderid           int64 
	Clordid              int64    
	Exchangeid           int       
	Exaccountid          int       
	Commodityid          int       
	Userid               int       
	Creatorloginid       int       
	Ordertype            int       
	Timeinforce          int64    
	Openclosetype        int       
	Buyselldirector      int       
	Orderprice           float64
	Orderquantity        float64
	Orderdate            int64    
	Lastdealprice        float64
	Lastdealquantity     float64
	Lastdealtime         int64    
	Currentdealquantity  float64
	Cancelquantity       float64
	Leavesquantity       float64
	Avgprice             float64
	Holdpositionid       int64    
	Limitorderid         int64    
	Otcorderid           string   
	Otcstopprofitid      string   
	Otcstoplossid        string   
	Updatetime           int64    
	Status               int       
}
