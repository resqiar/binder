package views

import (
	"binder/entities"
	"fmt"
	"strings"
)

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

func SendEditExtSuccessAlert(url string) string {
	return fmt.Sprintf(`
		<div role="alert" class="alert shadow-lg">
			<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
			<div>
				<h3 class="font-bold">Extension successfully edited!</h3>
				<div class="text-xs">You can continue editing or view details of the one you just edited.</div>
			</div>
			<a href="%s" class="btn btn-neutral btn-sm">See Details</a>
		</div>
		`,
		url,
	)
}

func SendSearchNotFound(message string) string {
	return fmt.Sprintf(`
		<div class="flex-cols mb-12 flex justify-center w-full flex-wrap gap-2 px-4 pb-20 md:flex-row lg:mt-4">
			<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
			  <path stroke-linecap="round" stroke-linejoin="round" d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 5.25h.008v.008H12v-.008Z" />
			</svg>

			<span>%s</span>
		</div>
		`,
		message,
	)
}

func SendSearchCard(exts []entities.Extension) string {
	var buf strings.Builder

	for _, val := range exts {
		buf.WriteString(fmt.Sprintf(`
		<div class="w-full md:w-auto">
			<a
				href="/ext/%s"
				class="image-full card w-full cursor-pointer shadow-md transition-all hover:-translate-y-1 hover:shadow-2xl lg:w-96"
			>
				<div class="card-body">
					<span class="badge badge-primary font-bold">%s</span>
					<h2 class="card-title">%s</h2>
					<p class="line-clamp-3">%s</p>
				</div>
			</a>
		</div>
		`,
			val.Slug,
			val.Slug,
			val.Title,
			val.Description.String,
		))
	}

	return buf.String()
}
