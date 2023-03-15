const crypto = require("crypto");

const keyMap = new Map();

function formatSha256(key) {
  if (keyMap.has(key)) {
    return keyMap.get(key);
  }

  const hash = crypto.createHash("sha256").update(key).digest("hex");
  const hashedKey = `0x${hash}`.toLowerCase();
  keyMap.set(key, hashedKey);

  return hashedKey;
}

module.exports.write = function write(key, value) {
  const hashedKey = formatSha256(key);

  process.env[`DB_${hashedKey}`] = value;

  if (typeof value === "bigint" || typeof value === "number") {
    value = `0x${value.toString(16)}`;
  } else if (typeof value === "boolean") {
    value = +value;
  }

  process.stdout.write(`DBW=${hashedKey}=${value}\n`);
};

module.exports.prune = function prune(key) {
  key = formatSha256(key); // convert to 32 byte key

  process.stdout.write(`DBP=${key}\n`);
  process.env[`DB_${key}`] = "";
};

module.exports.read = function read(key, defaultsTo) {
  const sha256key = formatSha256(key); // convert to 32 byte key

  const value = process.env[`DB_${sha256key}`];

  if (!value && defaultsTo === undefined) {
    process.stderr.write(
      `Database: Can't find key "DB_${key}" as "DB_${sha256key}" in env\n`
    );
    process.exit(1);
  }

  if (!value) {
    return defaultsTo;
  }

  return value;
};
