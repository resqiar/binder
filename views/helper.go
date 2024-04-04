package views

import "fmt"

func SendErrorAlert(message string) string {
	return fmt.Sprintf(`
		<div role="alert" class="alert alert-error mt-6">
			<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
			<span>%s</span>
		</div>
		`,
		message,
	)
}

func SendCreateExtSuccessAlert(url string) string {
	return fmt.Sprintf(`
		<div role="alert" class="alert shadow-lg">
			<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
			<div>
				<h3 class="font-bold">Extension successfully created!</h3>
				<div class="text-xs">You can create another extension or view details of the one you just created.</div>
			</div>
			<a href="%s" class="btn btn-neutral btn-sm">See Details</a>
		</div>
		`,
		url,
	)
}
