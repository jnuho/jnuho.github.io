grep -lR '"assets' articles | while read -r file; do
  sed -i 's/\"assets/\"..\/assets/g' "$file"
done
