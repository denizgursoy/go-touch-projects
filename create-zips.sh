

for f in *; do
    if [ -d "$f" ]; then
        echo "$d"
        zip -r "$d.zip" "$d/."
    fi
done