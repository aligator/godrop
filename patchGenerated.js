/**
 * This file is registered as hook for the graphql codegen to be run after file creation.
 * It patches the generated file to replace code which is not compatible
 * with the strict mode of typescript or has problems with the linter.
 */

const fs = require('fs')

let file = process.argv[2]

fs.readFile(file, 'utf8', function (err,data) {
    if (err) {
        return console.log(err);
    }

    // https://github.com/ticruz38/graphql-codegen-svelte-apollo/issues/22
    let result = data.replace("{ data: null, loading: true, error: null, networkStatus: 1, query: null }", 'undefined');

    // This code does not respect @typescript-eslint/no-unsafe-assignment
    result = result.replace("                  set({ ...v, query: q });", `                  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
                  set({ ...v, query: q });`)

    fs.writeFile(file, result, 'utf8', function (err) {
        if (err) return console.log(err);
    });
});