/**
 * @param {string[]} operations
 * @return {number}
 */
var calPoints = function(operations) {
    let angkaSimpan = [];
    let total = 0;

    for (let i = 0; i < operations.length; i++) {
        if (operations[i] == "C") {
            angkaSimpan.pop();

        } else if (operations[i] == "D") {
            let skorTerakhir = angkaSimpan[angkaSimpan.length - 1];
            angkaSimpan.push(2 * skorTerakhir);

        } else if (operations[i] == "+") {
            let skorTerakhir = angkaSimpan[angkaSimpan.length - 1];
            let skorKeduaTerakhir = angkaSimpan[angkaSimpan.length - 2];

            angkaSimpan.push(skorTerakhir + skorKeduaTerakhir);

        } else {
            angkaSimpan.push(Number(operations[i]));
        }
    }

    for (let i = 0; i < angkaSimpan.length; i++) {
        total += angkaSimpan[i];
    }

    return total;
};