#!/usr/bin/env nu

# Run go mod tidy in every lamu module (root + plugins).
# Usage: nu tidy.nu

def main [] {
  fd go.mod
  | lines
  | path dirname
  | par-each {|dir|
      let module = if ($dir | is-empty) or $dir == "." { "lamu" } else { $dir }
      print $"tidying ($module)"
      do -i { cd $dir; go mod tidy }
    }
}
