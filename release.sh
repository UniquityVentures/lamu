#!/usr/bin/env bash
set -euo pipefail

# Tag lamu and every plugin module at the same version.
# Usage: ./release.sh 0.4.14 [commit]
# Push:  git push origin v0.4.14 plugins/p_dashboard/v0.4.14 ...

version="${1:?Usage: ./release.sh <version> [commit]}"
version="${version#v}"
commit="${2:-HEAD}"
root_tag="v${version}"

plugins=(
	p_dashboard
	p_export
	p_filesystem
	p_google_genai
	p_livereloading
	p_otp
	p_pwa
	p_users
)

git tag -a "${root_tag}" -m "Release ${root_tag}" "${commit}"

plugin_tags=()
for plugin in "${plugins[@]}"; do
	tag="plugins/${plugin}/v${version}"
	git tag -a "${tag}" -m "Release ${root_tag}" "${commit}"
	plugin_tags+=("${tag}")
done

echo "Tagged ${root_tag} at ${commit}"
echo "Tagged plugins: ${plugin_tags[*]}"
echo
echo "Push with:"
echo "  git push origin ${root_tag} ${plugin_tags[*]}"
