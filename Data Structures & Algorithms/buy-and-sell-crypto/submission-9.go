func maxProfit(prices []int) int {
    maxP := 0
    price := math.MaxInt32

    for _, p := range(prices) {
        if int(p) < price {
            price = p
        } else {
            profit := int(p)-price
            maxP = max(maxP, profit)
        }
    }
    return maxP

}
