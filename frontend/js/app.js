document.addEventListener("DOMContentLoaded", () => {
    const bookForm = document.getElementById("book-form");

    async function loadBooks() {
        try {
            const response = await fetch("http://localhost:8081/api/books");

            if (!response.ok) {
                throw new Error("Failed to load books");
            }

            const books = await response.json();

            const booksList = document.getElementById("books-list");

            booksList.innerHTML = "";

            books.forEach((book) => {
                const row = document.createElement("tr");

                row.innerHTML = `
                    <td>${book.id}</td>
                    <td>${book.title}</td>
                    <td>${book.author}</td>
                `;

                booksList.appendChild(row);
            });

        } catch (error) {
            console.error("Error loading books:", error);
        }
    }

    bookForm.addEventListener("submit", async (event) => {
        event.preventDefault();

        const title = document.getElementById("title").value;
        const author = document.getElementById("author").value;
        const message = document.getElementById("form-message");

        const newBook = {
            title: title,
            author: author
        };

        try {
            const response = await fetch("http://localhost:8081/api/books", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(newBook)
            });

            const data = await response.json();

            if (!response.ok) {
                message.textContent = data.error || "Failed to create book";
                return;
            }

            message.textContent = "Book added successfully";

            bookForm.reset();

            await loadBooks();

        } catch (error) {
            console.error("Error creating book:", error);
            message.textContent = "Could not connect to the server";
        }
    });

    loadBooks();
});