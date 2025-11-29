# modrinth-sdk-go

## 0.2.0

### Minor Changes

- 0edf99a: Add changesets for automated version management and releases

  - Integrate changesets for version bumping and changelog generation
  - Add auto-release workflow that triggers on merged PRs with changesets
  - Add changeset check workflow to enforce changesets on PRs

### Patch Changes

- 4d3ce51: Remove duplicate release job from build workflow

  - Release is now handled exclusively by auto-release.yml with changesets
  - Prevents duplicate workflow runs on PR merge

## 0.1.0

### Minor Changes

- 63bb347: Set up CI/CD pipeline with GitHub Actions and Changesets for automated versioning and releases.
