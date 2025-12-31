package    lasagna

// TODO: define the 'OvenTime' constant// 待办事项：定义“OvenTime”常量
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.// RemainingOvenTime 根据已放入烤箱的“实际”分钟数返回剩余的分钟数。
 
func RemainingOvenTime(actual int) int {
    if actual>=40 {
    	return 0
    }
    return OvenTime - actual
    // TODO
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.// PreparationTime 函数根据千层面的层数计算准备所需的时间。
func   PreparationTime(numberOfLayers int) int {
	return numberOfLayers*2
    //	panic("PreparationTime not implemented")//panic//   恐慌("PreparationTime 未实现
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.// ElapsedTime 计算烹饪千层面所花费的时间。此时间包括准备时间和千层面在烤箱中烘烤的时间。
func   ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return numberOfLayers*2+actualMinutesInOven
    //	panic("ElapsedTime not implemented")//恐慌（“未实现 ElapsedTime”）
}
