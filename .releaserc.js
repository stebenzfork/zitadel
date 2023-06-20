module.exports = {
    branches: [
        { name: 'main' },
        { name: 'next' },
        { name: 'merge-eventstore-pipeline', prerelease: 'eventstore-performance'}
  ],
    plugins: [
      "@semantic-release/commit-analyzer",
      "@semantic-release/release-notes-generator",
      "@semantic-release/github"
    ]
};
