package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := map[string]int{}
	unitMap["quarter_of_a_dozen"] = 3
    unitMap["half_of_a_dozen"] = 6
    unitMap["dozen"] = 12
    unitMap["small_gross"] = 120
    unitMap["gross"] = 144
    unitMap["great_gross"] = 1728
    return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	billMap := map[string]int{}
    return billMap
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value , exist := units[unit]
    if !exist{
        return false
    }
    bill[item] = bill[item] + value
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value , exist := bill[item]
    if !exist{
        return false
    }
    value2 , exist2 := units[unit]
    if !exist2{
        return false
    }
    if(value < value2){
        return false
    }
    if(value ==value2){
        delete(bill,item)
        return true
    }
    bill[item] = bill[item] - units[unit]
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value,exist := bill[item]
    return value,exist
}
