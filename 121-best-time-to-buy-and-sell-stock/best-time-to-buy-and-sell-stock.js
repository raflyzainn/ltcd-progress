/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let lowest = prices[0];
    let maxProfit = 0;
    let currentProfit = 0;

    for (let i = 0; i < prices.length; i++) {
        if (prices[i] < lowest) {
            lowest = prices[i];
        } 

        currentProfit = prices[i] - lowest;
        
        if (currentProfit > maxProfit) {
            maxProfit = currentProfit;
        }
    }

    return maxProfit;
};