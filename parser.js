const ifJson = require('if-json')

function parser (args /* Array<string> */, popLast /* boolean */ = true) /* Object */{
  const argsArr = Object.keys(args)
  
  if (popLast) {
    // Remove Circular value from arguments
    argsArr.pop()
  }

  const result = {}

  argsArr.map(item => {
    const value = args[item]
    let [key, val = true] = value.split('=')

    const parsedVal = ifJson(val)
    val = parsedVal ? parsedVal : val

    result[key] = val
  })

  return result
}

module.exports = parser
