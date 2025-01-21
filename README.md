# Download Manager

This is a simple Download Manager project built in Go. It uses multithreading with `goroutine` and `channel` to efficiently download multiple files simultaneously. Users can dynamically input URLs for downloading files.

## Features
- Download files from multiple URLs concurrently.
- Dynamic user input for URLs.
- Saves downloaded files in the `downloads` directory.

---

## How to Use

### Prerequisites
- Make sure you have Go installed on your system.
- Clone or download the project to your local machine.

### Steps to Run

1. **Navigate to the project directory:**
   ```bash
   cd <project-directory>
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **Input the URLs of the files you want to download:**
   - After running the program, you'll be prompted to enter URLs.
   - Enter one URL per line and press Enter.
   - When you're done entering URLs, type `done` and press Enter to start the download.

   **Example input:**
   ```
   https://upload.wikimedia.org/wikipedia/commons/a/a6/Albert_Einstein_Head.jpg
   https://www.w3.org/TR/PNG/iso_8859-1.txt
   https://www.adobe.com/support/products/enterprise/knowledgecenter/media/c4611_sample_explain.pdf
   https://github.com/vmg/redcarpet/archive/refs/heads/master.zip
   done
   ```

4. **Monitor the progress:**
   - The program will display messages indicating the success or failure of each file download.

5. **Check the downloaded files:**
   - All downloaded files will be saved in the `downloads` directory inside the project folder.

---

## Example Output

```
Enter the URLs of files to download (type "done" when finished):
> https://upload.wikimedia.org/wikipedia/commons/a/a6/Albert_Einstein_Head.jpg
> https://www.w3.org/TR/PNG/iso_8859-1.txt
> https://www.adobe.com/support/products/enterprise/knowledgecenter/media/c4611_sample_explain.pdf
> https://github.com/vmg/redcarpet/archive/refs/heads/master.zip
> done
Successfully downloaded Albert_Einstein_Head.jpg
Successfully downloaded iso_8859-1.txt
Successfully downloaded c4611_sample_explain.pdf
Successfully downloaded master.zip
All downloads complete.
```

---

## Notes
- Ensure that the URLs you enter are valid and accessible.
- The application will create the `downloads` folder automatically if it doesn't exist.

Enjoy using the Download Manager! 🎉
