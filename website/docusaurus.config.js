// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Matrix-Veles',
  tagline: 'Protector of your Matrix-harvest!',
  url: 'https://veles.1in1.net',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'Unkn0wnCat',
  projectName: 'matrix-veles',
  i18n: {
    defaultLocale: "en",
    locales: ["en", "de"]
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: ({locale, versionDocsDirPath, docPath}) => {
            if (locale !== "en") {
              return `https://crowdin.com/project/matrix-veles/${locale}`;
            }
            return `https://github.com/Unkn0wnCat/matrix-veles/tree/main/website/${versionDocsDirPath}/${docPath}`;
          },
          showLastUpdateAuthor: true,
          showLastUpdateTime: true
        },
        /*blog: {
          showReadingTime: true,
          editUrl:
            'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        },*/
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Matrix-Veles',
        logo: {
          alt: 'Matrix-Veles Logo',
          src: 'img/logo.svg',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Documentation',
          },
          //{to: '/blog', label: 'Blog', position: 'left'},
          {
            type: 'localeDropdown',
            position: 'right',
          },
          {
            href: 'https://github.com/Unkn0wnCat/matrix-veles',
            label: 'GitHub',
            position: 'right',
          },
          {
            href: 'https://github.com/Unkn0wnCat/matrix-veles/releases/latest',
            label: 'Download',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: '/docs/intro',
              },
            ],
          },
          /*{
            title: 'Community',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/docusaurus',
              },
              {
                label: 'Discord',
                href: 'https://discordapp.com/invite/docusaurus',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/docusaurus',
              },
            ],
          },*/
          {
            title: 'More',
            items: [
              /*{
                label: 'Blog',
                to: '/blog',
              },*/
              {
                label: 'Imprint',
                href: 'https://kevink.dev/legal/about',
              },
              {
                label: 'Disclaimer',
                href: 'https://github.com/legal/disclaimer',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/facebook/docusaurus',
              },
            ],
          },
        ],
        copyright: `CC-BY-4.0 ${new Date().getFullYear()}, Matrix-Veles Contributors`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
      colorMode: {
        respectPrefersColorScheme: true
      },
      announcementBar: {
        id: 'alpha_software',
        content:
            'Veles is currently Alpha-Software - Some things might not be ready yet!',
        isCloseable: false,
      },
    }),
};

module.exports = config;
