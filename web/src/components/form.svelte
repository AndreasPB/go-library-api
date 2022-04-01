<script lang="ts">
	import type { Book } from "$lib/types"
</script>

<form>
	<!-- Grid -->
	<div class="grid">
		<!-- Markup example 1: input is inside label -->
		<label for="title">
			Title
			<input
				type="text"
				id="title"
				name="title"
				placeholder="The Lord of The Rings: The Fellowship of the Ring"
				required
			/>
		</label>

		<label for="author"
			>Author
			<input
				type="text"
				id="author"
				name="author"
				placeholder="J. R. R. Tolkien"
				required
			/>
		</label>
	</div>

	<div class="grid">
		<label for="id">
			ID
			<input type="number" id="id" name="id" placeholder="1234" required />
		</label>
		<label for="quantity"
			>Quantity
			<input
				type="number"
				id="quantity"
				name="quantity"
				placeholder="42"
				required
			/>
		</label>
		<button type="submit">Submit</button>
	</div>

	<!-- Button -->
</form>
