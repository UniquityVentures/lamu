#!/usr/bin/env nu

def main [tag: string] {
  nu ./tidy.nu
  fd go.mod | lines | path dirname | reverse | each { [$in $tag] | str join "/" | str trim -c "/" } | each {git tag $in}
  git push --tags
}
