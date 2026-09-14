document.addEventListener("DOMContentLoaded", () => {

    const booksGrid = document.getElementById("books-grid");
    const booksMessage = document.getElementById("books-message");

    async function loadBooks() {

        try {

            const response = await fetch(
                "http://localhost:8081/api/books"
            );

            if (!response.ok) {
                throw new Error("Failed to load books");
            }

            const books = await response.json();

            booksGrid.innerHTML = "";

            if (books.length === 0) {

                booksMessage.textContent =
                    "There are no books registered yet.";

                return;
            }

            booksMessage.textContent = "";

            books.forEach((book) => {

                const card = document.createElement("article");

                card.className = "book-card";

                card.innerHTML = `
                    <div class="book-cover">
                        <span>Book Cover</span>
                    </div>

                    <div class="book-card-content">

                        <h3>${book.title}</h3>

                        <p class="book-author">
                            ${book.author}
                        </p>

                        <span class="book-status">
                            Available
                        </span>

                        <button
                            type="button"
                            class="primary-button borrow-button"
                        >
                            Borrow
                        </button>

                    </div>
                `;

                booksGrid.appendChild(card);
            });

        } catch (error) {

            console.error("Error loading books:", error);

            booksMessage.textContent =
                "Could not connect to the server.";
        }
    }

    loadBooks();
});