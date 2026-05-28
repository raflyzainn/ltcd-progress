/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function(accounts) {
    const count = [];
    let currentValue = 0;

    for (let i = 0; i < accounts.length; i++) {
        count[i] = 0;
        for (let j = 0; j < accounts[i].length; j++) {
            count[i] = count[i] + accounts[i][j]
        }
    }

    let max = count[0];

    for (let k = 0; k < accounts.length; k++) {
        currentValue = count[k]
        if (currentValue > max) {
            max = currentValue
        } 
    }

    return max
};