// apollo.config.js
module.exports = {
  client: {
    service: {
      name: "personal-shopper",
      // URL to the GraphQL API
      url: "http://localhost:8080/query",
    },
    // Files processed by the extension
    includes: ["src/**/*.vue", "src/**/*.js"],
  },
};
