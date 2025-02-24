 export type Mail = {
	message_id: string
	date: Date
	datetime: Date
	from: string
	to: string
	bcc: string
	cc: string
	subject: string
	mime_version: string
	content_type: string
	content_transfer_encoding: string
	x_from: string
	x_to: string
	x_cc: string
	x_bcc: string
	x_folder: string
	x_origin: string
	x_file_name: string
	body: string
}

export type MailSummary =  {
  id:string
} & Pick<Mail,'from'|'to' | 'date' | 'subject'>
