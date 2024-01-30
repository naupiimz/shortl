<script lang="ts">
	import { enhance } from '$app/forms';
	import type { ActionData } from './$types';
	import { fade } from 'svelte/transition';
	import '../app.css';

	import { v4 as uuidv4 } from 'uuid';

	export let form: ActionData;

	let userId: string = uuidv4();

	let status: any = null;

	function copyToClipboard(): void {
		const inputElement = document.getElementById('result_url') as HTMLInputElement;
		const inputValue = (inputElement as HTMLInputElement).value;
		try {
			// Use the modern Clipboard API if available
			if (navigator.clipboard) {
				navigator.clipboard.writeText(inputValue).then(
					() => {
						status = true;
						setTimeout(() => {
							status = null;
						}, 5000);
					},
					(err) => {
						status = false;
						setTimeout(() => {
							status = null;
						}, 5000);
					}
				);
				return;
			}
		} catch (err) {
			console.error('Failed to copy text: ', err);
		}
	}
</script>

<div class="flex flex-col items-center justify-center gap-10 min-h-screen">
	<form class="flex flex-col items-center gap-3" method="POST" action="?/create" use:enhance>
		<h1 class="font-semibold text-4xl">shor.tl</h1>
		<div class="flex flex-row items-center gap-4">
			<input
				type="text"
				id="link"
				name="long_url"
				class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				placeholder="paste your link here!"
				required
			/>
			<input type="hidden" name="user_id" value={userId} />
			<button
				type="submit"
				class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
			>
				Generate
			</button>
		</div>
	</form>
	{#if form?.success}
		<section class="flex flex-row items-center gap-3 bg-green-400 py-10 px-10 rounded" in:fade>
			<input
				type="text"
				id="result_url"
				class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				value={form.result.short_url}
				disabled
			/>
			<button
				type="submit"
				class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
				on:click={copyToClipboard}
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="w-5 h-5"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M9 12h3.75M9 15h3.75M9 18h3.75m3 .75H18a2.25 2.25 0 0 0 2.25-2.25V6.108c0-1.135-.845-2.098-1.976-2.192a48.424 48.424 0 0 0-1.123-.08m-5.801 0c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75 2.25 2.25 0 0 0-.1-.664m-5.8 0A2.251 2.251 0 0 1 13.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m0 0H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V9.375c0-.621-.504-1.125-1.125-1.125H8.25ZM6.75 12h.008v.008H6.75V12Zm0 3h.008v.008H6.75V15Zm0 3h.008v.008H6.75V18Z"
					/>
				</svg>
			</button>

			<a
				href={form.result?.short_url}
				target="_blank"
				class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
				><svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="w-5 h-5"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9S9.515 3 12 3m0 0a8.997 8.997 0 0 1 7.843 4.582M12 3a8.997 8.997 0 0 0-7.843 4.582m15.686 0A11.953 11.953 0 0 1 12 10.5c-2.998 0-5.74-1.1-7.843-2.918m15.686 0A8.959 8.959 0 0 1 21 12c0 .778-.099 1.533-.284 2.253m0 0A17.919 17.919 0 0 1 12 16.5c-3.162 0-6.133-.815-8.716-2.247m0 0A9.015 9.015 0 0 1 3 12c0-1.605.42-3.113 1.157-4.418"
					/>
				</svg>
			</a>
		</section>
	{/if}
	{#if status == true}
		<div
			class="fixed bottom-16 right-8 text-gray-900 py-2 px-2 bg-green-400 rounded-md shadow-lg"
			in:fade={{ duration: 250 }}
			out:fade={{ duration: 250 }}
		>
			<div class="toast-header">
				<p class="font-semibold text-sm">Success!</p>
			</div>
			<div class="toast-body">
				<p class="text-sm text-gray-700">Copied to clipboard!</p>
			</div>
		</div>
	{:else if status == false}
		<div
			class="fixed bottom-16 right-8 text-gray-900 py-2 px-2 bg-red-400 rounded-md shadow-lg"
			in:fade={{ duration: 250 }}
			out:fade={{ duration: 250 }}
		>
			<div class="toast-header">
				<p class="font-semibold text-sm">Failed!</p>
			</div>
			<div class="toast-body">
				<p class="text-sm text-gray-700">Fail to copy to clipboard!</p>
			</div>
		</div>
	{/if}
</div>
