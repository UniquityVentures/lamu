#!/usr/bin/env nu

def main [tag: string] {
  fd go.mod | lines | path dirname | par-each { [$in $tag] | str join "/" | str trim -c "/" } | each {git tag $in}
  git push --tags
}
