# review

Review all uncommitted documentation changes for quality issues before committing.

## Instructions

You are tasked with reviewing all uncommitted changes in this documentation repository for quality issues.

### Step 1: Gather Changes

Run these commands in parallel to understand what changed:

1. `git status` - See all modified and untracked files
2. `git diff` - View unstaged changes
3. `git diff --cached` - View staged changes

### Step 2: Read Changed Files

Read the full content of each changed documentation file (`.mdx`, `.md`, `.json`) to perform a thorough review.

### Step 3: Review Checklist

For each changed file, check for the following issues:

#### Grammar & Spelling
- [ ] Spelling errors and typos
- [ ] Grammar mistakes
- [ ] Awkward phrasing or unclear sentences
- [ ] Inconsistent tense (prefer present tense for docs)
- [ ] Missing or incorrect punctuation

#### Links & References
- [ ] Broken internal links (paths to other `.mdx` files)
- [ ] Broken anchor links (`#section-name`)
- [ ] Broken snippet references (`<Snippet file="..." />`)
- [ ] Broken image references
- [ ] External links that look malformed

To verify internal links exist, use `ls` or glob to check if referenced files exist.

#### Markdown & Formatting
- [ ] Proper heading hierarchy (don't skip levels, e.g., h1 → h3)
- [ ] Consistent list formatting (bullets vs numbers)
- [ ] Code blocks have language tags (```python, ```bash, etc.)
- [ ] Proper MDX component syntax
- [ ] Unclosed tags or brackets

#### Content Quality
- [ ] Incomplete sentences or TODO markers left in
- [ ] Placeholder text (lorem ipsum, "TBD", "FIXME")
- [ ] Outdated information that contradicts other docs
- [ ] Missing context that would confuse readers
- [ ] Steps that are out of order or missing

#### Frontmatter (if applicable)
- [ ] Valid YAML syntax in frontmatter
- [ ] Required fields present (title, description)
- [ ] Description is meaningful (not placeholder)

#### Consistency
- [ ] Terminology matches the rest of the docs (check `snippets/definitions/` for canonical terms)
- [ ] UI element names match actual product
- [ ] Command examples are accurate

### Step 4: Present Findings

Present your findings in this format:

```
## Review Summary

**Files reviewed:** N files
**Issues found:** N issues

---

### filename.mdx

#### 🔴 Errors (must fix)
- Line 23: Broken link to `/docs/missing-page.mdx` - file does not exist
- Line 45: Unclosed code block

#### 🟡 Warnings (should fix)
- Line 12: Grammar issue - "This allows to configure" → "This allows you to configure"
- Line 34: Code block missing language tag

#### 🔵 Suggestions (optional)
- Line 56: Consider simplifying "utilize" → "use"

---

### another-file.mdx
...
```

### Step 5: Offer to Fix

After presenting findings, ask:

**"Would you like me to fix any of these issues? Reply with:**
- **'all'** - Fix all errors and warnings
- **'errors'** - Fix only errors
- **specific numbers** - e.g., 'fix 1, 3, 5'
- **'no'** - Don't fix anything"

### Step 6: Apply Fixes (If Requested)

If the user requests fixes:
1. Make the requested changes using file edit tools
2. Show a summary of what was fixed
3. Do NOT commit - leave that to the user

### Important Rules

- **Be specific** - Always include line numbers and exact text
- **Verify links** - Actually check if referenced files exist before reporting broken links
- **Don't be pedantic** - Focus on real issues, not style preferences
- **Respect conventions** - This repo may have intentional style choices
- **Read context** - Some "issues" may be intentional (e.g., placeholder in a template)
- **Prioritize** - Errors > Warnings > Suggestions
