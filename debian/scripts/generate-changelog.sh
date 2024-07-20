#!/bin/sh
set -ex

package=$(grep '^Source:' debian/control | awk '{print $2}')

orig_commit=$(git rev-parse HEAD)

tags=$(git tag --sort=creatordate)

prev_tag=""
for tag in $tags; do
  git checkout $tag

  new_version="$(echo $tag | tr -d 'v')-1"

  export FAKETIME=$(git show -s --format=%aI $tag | sed 's/T/ /; s/.\{6\}$//')
  
  if [ -n "$prev_tag" ]; then
    LD_PRELOAD="/usr/lib/$(uname -m)-linux-gnu/faketime/libfaketime.so.1" \
      gbp dch --ignore-branch --release --new-version="$new_version" --since=$prev_tag --spawn-editor=never
  else
    cat <<EOF > debian/changelog
$package (0.0.0-1) UNRELEASED; urgency=medium

 -- $DEBFULLNAME <$DEBEMAIL>  Thu, 01 Jan 1970 00:00:00 +0000
EOF

    LD_PRELOAD="/usr/lib/$(uname -m)-linux-gnu/faketime/libfaketime.so.1" \
      gbp dch --ignore-branch --release --new-version=$new_version --spawn-editor=never
  fi

  prev_tag=$tag
done

git checkout $orig_commit