module.exports.write = function write(key, value) {
  process.env[`DB_${key}`] = value;

  if (typeof value === "bigint" || typeof value === "number") {
    value = `0x${value.toString(16)}`;
  } else if (typeof value === "boolean") {
    value = +value;
  }

  process.stdout.write(`DBW=${key}=${value}\n`);
};

module.exports.prune = function prune(key) {
  process.stdout.write(`DBP=${key}\n`);
  process.env[`DB_${key}`] = "";
};

module.exports.read = function read(key, defaultsTo) {
  const value = process.env[`DB_${key}`];

  if (!value && defaultsTo === undefined) {
    process.stderr.write(`Database: Can't find key "DB_${key}" in env\n`);
    process.exit(1);
  }

  if (!value) {
    return defaultsTo;
  }

  return value;
};
