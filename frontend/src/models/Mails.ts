 export interface MailSummary{
  id:string;
  from:string;
  to:string;
  date:Date
  subject:string;
  x_from: string;
	x_to: string;
 }

 export interface Mail extends MailSummary {
	message_id: string;
	datetime: Date;
	bcc: string;
	cc: string;
	mime_version: string;
	content_type: string;
	content_transfer_encoding: string;
	x_cc: string;
	x_bcc: string;
	x_folder: string;
	x_origin: string;
	x_file_name: string;
	body: string;
}
