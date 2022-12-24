find . -name "*.go"|while read fname; do
  go fmt $fname
done

exit 0

