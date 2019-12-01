/*
const commands = {
    next:"",
    prev:"",
    inc:"",
    dec:"",
    write:"",
    open:"",
    close:""
} 
*/

const commands = {
  next: "😩",
  prev: "😄",
  inc: "😏",
  dec: "😆",
  write: "😡",
  open: "😍",
  close: "😴"
};

const maxStep = 10000;
const maxMemory = 100;

const start = input => {
  const inputCommands = Array.from(input.trim().replace(/\r?\n/g, ""));
  let memory = Array.apply(null, Array(maxMemory)).map(function() {
    return 0;
  });
  let turn = 0;
  let pointer = 0;
  let output = "";
  let step = 0;
  run: while (turn < inputCommands.length) {
    if (step > maxStep) {
      break run;
    } else {
      step++;
    }
    switch (inputCommands[turn]) {
      case commands.next: {
        pointer++;
        break;
      }
      case commands.prev: {
        pointer--;
        break;
      }
      case commands.inc: {
        memory[pointer]++;
        break;
      }
      case commands.dec: {
        memory[pointer]--;
        break;
      }
      case commands.write: {
        output += String.fromCharCode(memory[pointer]);
        break;
      }
      case commands.open: {
        if (memory[pointer] == 0) {
          let depth = 1;
          while (depth > 0) {
            turn++;
            let nCommand = inputCommands[turn];
            if (nCommand === commands.open) {
              depth++;
            }
            if (nCommand === commands.close) {
              depth--;
            }
          }
        }
        break;
      }
      case commands.close: {
        let depth = 1;
        while (depth > 0) {
          turn--;
          const nCommand = inputCommands[turn];
          if (nCommand === commands.open) {
            depth--;
          }
          if (nCommand === commands.close) {
            depth++;
          }
        }
        turn--;
        break;
      }
      default: {
        output = "errorToken:" + turn;
        break run;
      }
    }
    turn++;
  }
  return output;
};

export default start;
