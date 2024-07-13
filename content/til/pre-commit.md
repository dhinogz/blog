+++
title = 'TIL: Git Pre-Commits'
date = 2024-07-12T14:00:46-06:00
draft = false
+++

As I was configuring the deployment of this blog, I wanted to see how I could automate the build and publish process. I knew using Github Actions was a common approach. Here's what I envisioned in my head:

![First Attempt CI/CD](/images/first-attempt-cicd-blog.png)

Implementing this approach caused me a bit of trouble with managing the public folder's Git history and was a headache to get working properly. Best to toss it and start over again. Enter Git Hooks.

There are different types of [Git Hooks](https://git-scm.com/docs/githooks), but the one I that fitted my needs was the pre-commit hook. Here's how I reimagined my approach.

![Pre Commits CI/CD](/images/pre-commit-cicd-blog.png)

This way, I avoid any git version conflict and never miss a build when pushing.

This is how my pre-commit hook looks like at .git/hooks/pre-commit

```bash
#!/bin/sh

hugo --minify
git add public/
echo "Generated public static files"
```
