# vocabulary action

A Github Action to create and update lists of vocabulary.

## Usage

> **Note**
> The Action has to be used in a Job that runs on a UNIX system (e.g. `ubuntu-latest`).

Example workflow:

```yaml
name: Add Vocabulary

on:
  workflow_dispatch:
    inputs:
      language:
        description: Language of the vocabulary list
        required: true
        type: string
      vocabulary:
        description: Vocabulary list, one word per line
        required: true
        type: string

jobs:
  vocabulary:
    runs-on: ubuntu-latest

    # permission required for committing to the repository
    permissions:
      contents: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3
        with:
          ref: ${{ github.head_ref }}

      - name: Add vocabulary
        id: vocabulary
        uses: leonhfr/vocabulary-action@release
        with:
          language: ${{ inputs.language }}
          vocabulary: ${{ inputs.vocabulary }}

      - name: Commit
        uses: stefanzweifel/git-auto-commit-action@v4
        with:
          file_pattern: ${{ steps.vocabulary.outputs.directory }}
          commit_message: "Add vocabulary: ${{ steps.vocabulary.outputs.summary }}"
```

## Config

The action looks for a config file named `vocabulary.yml` in the root of the repository:

```yaml
languages:
  ca: catalan/vocabulary
  de: german/vocabulary
  en: english/vocabulary
  es: spanish/vocabulary
  default: todo/vocabulary # not required
```

## Inputs

The action requires two inputs:

- `language`: The language of the vocabulary list. Will be used to lookup the target directory in the config file.
- `vocabulary`: Vocabulary list, one word per file.

## Outputs

You can use the outputs in other actions:

- `directory`: The repository directory modified by the action.
- `summary`: Summary of added vocabulary, intended to be used in a commit message.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/leonhfr/vocabulary-action/blob/master/LICENSE) file for details.
