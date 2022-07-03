const {
    gray, blueBright, red,
    yellowBright, dim, whiteBright
} = require("colorette");

const Print = (...msg) => process.stdout.writable ? process.stdout.write(msg + "\n") : console.log(msg);

const LoopThru = (array) => {
    for (const element of array) {
        Print(`${gray("[")}${blueBright("FOR-OF")} ${dim(yellowBright("LOOP"))}${gray("]")} Item: ${element}\n`);
    }

    array.forEach((item) => {
        Print(`${gray("[")}${blueBright("FOR-EACH")} ${dim(yellowBright("LOOP"))}${gray("]")} Item: ${item}`);
    });

    const m = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    Print(`\n${m.reverse()}\n`);
};

const TimeDevider = (timer_one, timer_two) => {
    return Math.abs(timer_one - timer_two);
};

const main = () => {
    const timer_one = Date.now();

    LoopThru(
        [
            "This",
            "is",
            "that",
            "and",
            "that",
            "is",
            "this",
            "and",
            "I",
            "am",
            "not",
            "possible"
        ]
    );

    const timer_two = Date.now();

    return Print(`${gray("[")}${red("DEVISION")}${gray("]")} ${TimeDevider(timer_one, timer_two)}ms`);
};

main();