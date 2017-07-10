const program = require('commander')
const ifJson = require('if-json')
const npmPkg = require('./package.json')

program
  .version('0.0.0')
  .action(function () {
    const args = Object.keys(arguments)
    // Remove Circular value from arguments
    args.pop()

    const result = {}
    args.map(item => {
      const value = arguments[item]
      let [key, val = true] = value.split('=')

      const parsedVal = ifJson(val)
      val = parsedVal ? parsedVal : val

      result[key] = val
    })

    console.log(JSON.stringify(result, null, 2))
  })
  .parse(process.argv)
