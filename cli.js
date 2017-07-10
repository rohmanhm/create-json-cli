const program = require('commander')
const npmPkg = require('./package.json')
const parser = require('./parser')

program
  .version(npmPkg.version)
  .action(function () {
    const result = parser(arguments)

    console.log(JSON.stringify(result, null, 2))
  })
  .parse(process.argv)
