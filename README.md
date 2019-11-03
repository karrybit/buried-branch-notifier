# buried-branch-notifier

[![Actions Status](https://github.com/makanai5610/buried-branch-notifier/workflows/Build/badge.svg)](https://github.com/makanai5610/buried-branch-notifier/actions)
[![Actions Status](https://github.com/makanai5610/buried-branch-notifier/workflows/Test/badge.svg)](https://github.com/makanai5610/buried-branch-notifier/actions)
[![](https://tokei.rs/b1/github/makanai5610/buried-branch-notifier?category=files)](https://github.com/makanai5610/buried-branch-notifier)
[![](https://tokei.rs/b1/github/makanai5610/buried-branch-notifier?category=lines)](https://github.com/makanai5610/buried-branch-notifier)

This tool notifies a list of buried branches to slack.

## How it is done
First, execute `git branch -r` to get a list of all branches.
```
$ git branch -r
```

Second, execute the following command to get last commit informations in JSON format.
```
$ git log -n 1 <branch_name> --pretty=format:{"branch_name":"<branch_name>","commiter_name":"%aN","last_commit_date":"%ad"} --date=iso-strict
```

Third, notify the decoding result of JSON to slack.
