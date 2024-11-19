# fee-free-ticketing
Why pay event platform fees when you can run your own micro-service instead?

## Why?
Using a managed event-ticketing site makes things easy - for a fee. These fees may seem inocuous, but they add up.

| Event Platform | Fees (in US) |
|  -----         | ---- |
| Eventbrite | 3.7% of Ticket Price + $1.79 per ticket + payment processing fees |
| TicketLeap | 2% of Ticket Price + $1.00 per ticket + payment processing fees |

Even for a modest event with ~$20 tickets and ~100 attendees, this equates to $100s of dollars in ticket fees - exclusive of payment processing costs!

### Total Cost of an event:
- Assuming use of the GCP free tier, only Stripe credit-card processing fees - that's it!

## Getting started
#### Disclaimer
This repo is meant to be extremely lightweight - you will not find enterprise CI/CD pipelines, IaC-managed infrastructure, or observability here.

This approach is likely only worthwhile from a time and money perspective if you already have the corporate/LLC entity established to be able to use Stripe and Sendgrid.

### Tech Stack:
- Stripe Account (handling payments)
- Sendgrid Account (email notifications)
- Google Cloud Run (compute, coordinate confirmation email firing)

### Creating Your Event
1. Create a payment link in Stripe for your event.
2. Create a webhook in Stripe to forward `checkout.session.completed` events to an arbitrary domain name.
3. Create a Sendgrid email template with fields for `first_name`, `ticket_count` and `total_cost`.
4. Create secrets in GCP for:
   - `STRIPE_WEBHOOK_KEY`
   - `SENDGRID_API_KEY`
   - `SENDGRIDEMAILTEMPLATEID`
The default compute engine service account must be granted access to read each of these secret values.

5. Fork this repository and deploy to your Google Cloud account.
Update env vars and values specific for your GCP project, email addresses, etc. 
   1. `just set-project`. Set build and deployment to be in the right GCP project. You may need to authenticate your `gcloud` CLI tool with `gcloud auth login` first.
   2. `just create-artifact-registry-repo`. Create a GCP artifact registry repo for the services image.
   3. `just build-and-push-image`. Build the Docker image and push it to your remote repository.
   4. `just deploy-cloud-run`. Creates a cloud-run instance to serve traffic. Go to the cloud console and grab the API endpoint that GCP configures.
6. Update the Stripe webhook event to send requests to the service's Google Cloud Run endpoint + `/api/v1`.
7. Notify attendees to register for their fee-free event (expect for credit card fees, impossible to avoid those).

Having now saved some dollar amount on platform fees, realize that sometimes it is worth paying someone to manage a service for you and use Eventbrite etc. for subsequent events 😅.
