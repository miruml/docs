# Copying a Folder Inside a Cloudflare R2 Bucket

This guide explains how to copy a “folder” (object prefix) inside a Cloudflare R2 bucket so multiple versions of assets (e.g. docs images) can exist at the same time.

Example use case:
- Source: `docs/images`
- Destination: `docs/v01/images`

---

### Install rclone

```bash
# macOS
brew install rclone

# Ubuntu / Debian
sudo apt install rclone
```

---


## Access Key ID and Secret

Go to the R2 object storage overview in cloudflare. 

Under **Account Details** on the right of the page is a card that says API Tokens. Click the **Manage** button on the right of the card.

Next create a **User API Token** with:
- **Object Read & Write** Permissions
- Only the `public-assets` bucket
- TTL of 24 hours

After creation, you should see the access key id and secret access key fields available to copy.

## One-Time Setup (Create the `r2` Remote)

This only needs to be done once per machine.

```bash
rclone config
```

Follow these prompts:

1. New remote → `n`
2. Name → `r2`
3. Storage → `s3`
4. Provider → `Cloudflare`
5. Access Key ID → `<YOUR_R2_ACCESS_KEY_ID>`
6. Secret Access Key → `<YOUR_R2_SECRET_ACCESS_KEY>`
7. Region → *(leave empty)*
8. Endpoint → `https://<ACCOUNT_ID>.r2.cloudflarestorage.com`
11. Advanced config → `n`
12. Save → `y`

---

## Copy a Folder (Prefix) Inside the Bucket

### Basic copy

```bash
rclone copy r2:public-assets/docs/images r2:public-assets/docs/v01/images
```

This:
- Copies all objects under `docs/images/`
- Preserves subfolders and filenames
- Leaves the original files untouched
- Can be safely re-run

---

### Dry run

```bash
rclone copy r2:public-assets/docs/images r2:public-assets/docs/v01/images --dry-run
```

---

### Show progress

```bash
rclone copy r2:public-assets/docs/images r2:public-assets/docs/v01/images --progress
```

---

## Delete a Folder (Prefix) Inside the Bucket

Since R2 doesn't have real folders, deleting a "folder" means deleting all objects with that prefix.

### Basic delete (recommended)

```bash
rclone delete r2:public-assets/docs/v01/images
```

This:
- Deletes all objects under `docs/v01/images/`
- Includes all subfolders and files
- ⚠️ **This action cannot be undone**

---

### Dry run (recommended first)

```bash
rclone delete r2:public-assets/docs/v01/images --dry-run
```

This shows what would be deleted without actually deleting anything.

---

### Troubleshooting: "Access Denied" error with `purge`

If you see an error like:
```
ERROR : S3 bucket public-assets path docs/images: Failed to read versioning status, assuming unversioned: AccessDenied
```

This happens because `rclone purge` tries to check bucket versioning status, which R2 doesn't support the same way as AWS S3. Use `rclone delete` instead (shown above), which works reliably with R2.

Alternatively, if you need to use `purge`, you can disable the versioning check by adding the `--s3-disable-checksum` flag:

```bash
rclone purge r2:public-assets/docs/v01/images --s3-disable-checksum
```

---

## Notes

- Cloudflare R2 does not have real folders — these are object key prefixes.
- Copying folders duplicates storage (both versions coexist).
- This is a server-side copy (no local download).

---

## Optional: Exact Mirroring (Use With Caution)

```bash
rclone sync r2:public-assets/docs/images r2:public-assets/docs/v01/images
```

⚠️ This will delete any files in `docs/v01/images` that don’t exist in `docs/images`.

---