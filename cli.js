#!/usr/bin/env node

var kiwi = require('./kiwi');
var fs = require('fs');

var usage = [
  '',
  'Usage: kiwic [OPTIONS]',
  '',
  'Options:',
  '',
  '  --help                Print this message.',
  '  --schema [PATH]       The schema file to use.',
  '  --js [PATH]           Generate JavaScript code.',
  '  --ts [PATH]           Generate TypeScript type definitions.',
  '  --cpp [PATH]          Generate C++ code (tree style).',
  '  --callback-cpp [PATH] Generate C++ code (callback style).',
  '  --skew [PATH]         Generate Skew code.',
  '  --go [PATH]           Generate Go code (tree style).',
  '  --binary [PATH]       Encode the schema as a binary blob.',
  '  --root-type [NAME]    Set the root type for JSON.',
  '  --to-json [PATH]      Convert a binary file to JSON.',
  '  --from-json [PATH]    Convert a JSON file to binary.',
  '',
  'Examples:',
  '',
  '  kiwic --schema test.kiwi --js test.h',
  '  kiwic --schema test.kiwi --cpp test.h',
  '  kiwic --schema test.kiwi --skew test.sk',
  '  kiwic --schema test.kiwi --binary test.bkiwi',
  '  kiwic --schema test.kiwi --root-type Test --from-json buffer.json',
  '  kiwic --schema test.kiwi --root-type Test --to-json buffer.bin',
  '',
].join('\n');

var main = exports.main = function(args) {
  var flags = {
    '--schema': null,
    '--js': null,
    '--ts': null,
    '--cpp': null,
    '--callback-cpp': null,
    '--skew': null,
    '--go': null,
    '--binary': null,
    '--root-type': null,
    '--to-json': null,
    '--from-json': null,
  };

  // Parse flags
  for (var i = 0; i < args.length; i++) {
    var arg = args[i];

    if (arg === '-h' || arg === '--help' || arg[0] !== '-') {
      console.log(usage);
      return 1;
    }

    else if (arg in flags) {
      if (i + 1 === args.length) {
        throw new Error('Missing value for ' + JSON.stringify(arg) + ' (use "--help" for usage)');
      }
      flags[arg] = args[++i];
    }

    else {
      throw new Error('Unknown flag ' + JSON.stringify(arg) + ' (use "--help" for usage)');
    }
  }

  // Must have a schema
  if (flags['--schema'] === null) {
    console.log(usage);
    return 1;
  }

  // Try loading the schema
  var content = fs.readFileSync(flags['--schema'], 'utf8');

  // Try parsing the schema, pretty-print errors on failure
  try {
    var schema = kiwi.compileSchema(content);
  } catch (e) {
    if (e && e.message && 'line' in e || 'column' in e) {
      e.message =
        flags['--schema'] + ':' + e.line + ':' + e.column + ': error: ' + e.message + '\n' +
        content.split('\n')[e.line - 1] + '\n' + new Array(e.column).join(' ') + '^';
    }
    throw e;
  }

  // Validate the root type
  var rootType = flags['--root-type'];
  if (rootType !== null && !(('encode' + rootType) in schema && ('decode' + rootType) in schema)) {
    throw new Error('Invalid root type: ' + JSON.stringify(rootType));
  }

  // Generate JavaScript code
  if (flags['--js'] !== null) {
    fs.writeFileSync(flags['--js'], kiwi.compileSchemaJS(content));
  }

  // Generate JavaScript code
  if (flags['--ts'] !== null) {
    fs.writeFileSync(flags['--ts'], kiwi.compileSchemaTypeScript(content));
  }

  // Generate C++ code
  if (flags['--cpp'] !== null) {
    fs.writeFileSync(flags['--cpp'], kiwi.compileSchemaCPP(content));
  }
  if (flags['--callback-cpp'] !== null) {
    fs.writeFileSync(flags['--callback-cpp'], kiwi.compileSchemaCallbackCPP(content));
  }

  // Generate Skew code
  if (flags['--skew'] !== null) {
    fs.writeFileSync(flags['--skew'], kiwi.compileSchemaSkew(content));
  }

  // Generate Go code
  if (flags['--go'] !== null) {
    fs.writeFileSync(flags['--go'], kiwi.compileSchemaGo(content));
  }

  // Generate a binary schema file
  if (flags['--binary'] !== null) {
    fs.writeFileSync(flags['--binary'], Buffer(kiwi.encodeBinarySchema(content)));
  }

  // Convert a binary file to JSON
  if (flags['--to-json'] !== null) {
    if (rootType === null) {
      throw new Error('Missing flag --root-type when using --to-json');
    }
    fs.writeFileSync(flags['--to-json'] + '.json', JSON.stringify(schema['decode' + rootType](new Uint8Array(fs.readFileSync(flags['--to-json']))), null, 2) + '\n');
  }

  // Convert a JSON file to binary
  if (flags['--from-json'] !== null) {
    if (rootType === null) {
      throw new Error('Missing flag --root-type when using --from-json');
    }
    fs.writeFileSync(flags['--from-json'] + '.bin', Buffer(schema['encode' + rootType](JSON.parse(fs.readFileSync(flags['--from-json'], 'utf8')))));
  }

  return 0;
};

if (require.main === module) {
  try {
    process.exit(main(process.argv.slice(2)));
  } catch (e) {
    process.stderr.write((e && e.message || e) + '\n');
    process.exit(1);
  }
}
