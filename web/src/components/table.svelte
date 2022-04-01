<script lang="ts">
	import type { Book } from "$lib/types"

	const updateBooks = async () => {
		return await fetch("http://0.0.0.0:8080/books").then((res) => res.json())
	}

	let books: Promise<Book[]> = updateBooks()

	const returnBook = async (id: string) => {
		await fetch(`http://0.0.0.0:8080/return?id=${id}`, {
			method: "PATCH"
		}).then((res) => res.json())

		// Refreshes table
		books = updateBooks()
	}

	const checkoutBook = async (id: string) => {
		const res = await fetch(`http://0.0.0.0:8080/checkout?id=${id}`, {
			method: "PATCH"
		})

		if (res.status === 200) {
			// Refreshes table
			books = updateBooks()
		} else {
			alert("Book is out of stock")
		}
	}
</script>

{#await books}
	<p>Loading...</p>
{:then books}
	<table>
		<thead>
			<tr>
				<th scope="col">#</th>
				<th scope="col">Title</th>
				<th scope="col">Author</th>
				<th scope="col">Quantity</th>
				<th scope="col">Checkout</th>
				<th scope="col">Return</th>
			</tr>
		</thead>
		<tbody>
			{#each books as book}
				<tr>
					<td>{book.id}</td>
					<td>{book.title}</td>
					<td>{book.author}</td>
					<td>{book.quantity}</td>
					{#if book.quantity > 0}
						<td>
							<button on:click={(_) => checkoutBook(book.id)}>Checkout</button>
						</td>
					{:else}
						<td>
							<button disabled>Checkout</button>
						</td>
					{/if}
					<td>
						<button on:click={(_) => returnBook(book.id)}>Return</button>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{:catch error}
	<p>{error}</p>
{/await}
