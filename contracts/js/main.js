const callintrpr = require("./utils/CallInterpreter/callInterpreter");

const QAN20 = require("./QAN20/qan20");

const qan20 = new QAN20("Example", "XMPL", 18n);

callintrpr.interpret(qan20, process.argv.slice(2));
