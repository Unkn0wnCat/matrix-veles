module.exports = {
    createOldCatalogs: true,
    lexers: {
        js: ['JsxLexer'],
        default: ['JavascriptLexer'],
        ts: ['JavascriptLexer'],
        jsx: ['JsxLexer'],
        tsx: ['JsxLexer'],
    },
    skipDefaultValues: (locale, ns) => {return locale !== "en"},
    locales: ['en', 'de'],
    output: 'public/locales/$LOCALE/$NAMESPACE.json',
    input: [
        'src/**/*.tsx',
        'src/*.tsx',
    ],
}