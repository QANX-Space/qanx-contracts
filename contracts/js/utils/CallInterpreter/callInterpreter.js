module.exports.interpret = function interpret(contract, args) {
  let [methodName, ...methodArguments] = args;

  methodArguments = convertArgs(methodArguments);

  if (!contract[methodName]) {
    process.stderr.write(
      `CallInterpreter: Method "${methodName}" does not exist\n`
    );
    process.exit(1);
  }

  if (typeof contract[methodName] !== "function") {
    process.stderr.write(`CallInterpreter: "${methodName}" is not a method\n`);
    process.exit(1);
  }

  if (contract[methodName].length !== methodArguments.length) {
    process.stderr.write(
      `CallInterpreter: Expected ${contract[methodName].length} arguments received ${methodArguments.length}\n`
    );
    process.exit(1);
  }

  const result = contract[methodName](...methodArguments);

  if (result) {
    if (Array.isArray(result)) {
      process.stdout.write(`OUT=${result.map(convertToString).join(" ")}`);
    } else {
      process.stdout.write(`OUT=${convertToString(result)}`);
    }
  }

  process.exit(0);
};

function convertArgs(args) {
  // pre fill with strings
  const converted = [...args];

  for (let i = 0; i < converted.length; i++) {
    const arg = converted[i];

    // 0/1 booleans
    if (/([01])/g.test(arg)) {
      convertArgs[i] = Boolean(arg);
    }

    // True/False/true/false booleans
    else if (/true|false/gi.test(arg)) {
      convertArgs[i] = /true/gi.test(arg);
    }

    // Hex number
    else if (/0x([0-9]|[a-f])+/gi.test(arg)) {
      convertArgs[i] = BigInt(arg);
    }

    // Normal numbers
    else if (/^\d+$/.test(arg)) {
      convertArgs[i] = BigInt(arg);
    }
  }

  return converted;
}

function convertToString(val) {
  switch (typeof val) {
    default:
      return val.toString();
    case "bigint":
    case "number":
      return val.toString(10);
    case "boolean":
      return (+val).toString(10);
  }
}
