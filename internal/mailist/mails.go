package mailist

import (
	"bytes"
	"html/template"
	"time"
)

type EmailData struct {
	UserName       string
	UserScore      int
	ReadinessLevel string
	RatingTitle    string
	Insight        string
	ActionSteps    []string
	Color          string
	Emoji          string
	PrimaryColor   string
	SecondaryColor string
	CurrentYear    int
}

func GenerateAIReadinessEmail(userName string, userScore int) (string, error) {
	data := EmailData{
		UserName:       userName,
		UserScore:      userScore,
		PrimaryColor:   "#0b3937",
		SecondaryColor: "#a9fa60",
		CurrentYear:    time.Now().Year(),
	}

	if userScore >= 80 {
		data.ReadinessLevel = "AI Leader"
		data.RatingTitle = "Highly Ready"
		data.Insight = "Your business is primed to scale AI as a core capability. You have strong strategy, infrastructure, and culture. Focus on maximizing your competitive advantage."
		data.ActionSteps = []string{
			"Scale AI across multiple functions for maximum impact and integration.",
			"Explore advanced applications (Generative AI, predictive models) to drive innovation.",
			"Develop an enterprise-wide AI playbook and robust governance model.",
			"Build a long-term AI roadmap aligned to innovation or product strategy.",
		}
		data.Color = "#4CAF50"
		data.Emoji = "✅"
	} else if userScore >= 60 {
		data.ReadinessLevel = "AI Explorer"
		data.RatingTitle = "Moderately Ready"
		data.Insight = "You have a solid foundation and cultural openness to begin embedding AI more deeply into operations, but some gaps remain, particularly in scaling."
		data.ActionSteps = []string{
			"Scale successful pilots across departments to realize measurable ROI.",
			"Address specific weak spots (e.g., governance, infrastructure, staff skills) revealed in the assessment.",
			"Prioritize 2–3 high-value AI use cases with measurable returns and clear ownership.",
			"Build a structured change management approach to ensure organization-wide adoption.",
		}
		data.Color = "#2196F3"
		data.Emoji = "🚀"
	} else if userScore >= 40 {
		data.ReadinessLevel = "AI Beginner"
		data.RatingTitle = "Low Readiness"
		data.Insight = "Interest exists, and you show potential, but critical foundations are missing, particularly in integrated strategy, technology, and data governance."
		data.ActionSteps = []string{
			"Improve data collection and governance; focus on data quality and accessibility.",
			"Start with low-risk pilots to build confidence and gather internal case studies.",
			"Invest strategically in staff training and AI literacy across all departments.",
			"Assign an “AI Champion” internally to drive momentum and coordinate initiatives.",
			"Begin documenting SOPs for future AI-powered workflows.",
		}
		data.Color = "#FF9800"
		data.Emoji = "🟡"
	} else {
		data.ReadinessLevel = "AI Unprepared"
		data.RatingTitle = "Very Low Readiness"
		data.Insight = "You are just beginning your AI journey. Focus first on leadership buy-in, data quality, and basic digital transformation before advanced AI integration."
		data.ActionSteps = []string{
			"Develop leadership alignment on the strategic value of AI.",
			"Run internal workshops on AI potential and risks to build foundational awareness.",
			"Train employees on AI Foundations and basic tools (e.g., specific Copilots).",
			"Audit your existing digital systems and address critical infrastructure gaps.",
			"Identify one non-technical team to experiment with a simple AI tool (e.g., HR, marketing).",
		}
		data.Color = "#F44336"
		data.Emoji = "🛑"
	}

	const tmpl = `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <style>
            body { font-family: 'Arial', sans-serif; background-color: #f7f9fc; margin: 0; padding: 0; }
            .container { max-width: 600px; margin: 30px auto; background-color: #ffffff; border-radius: 12px; border: 1px solid #e0e0e0; overflow: hidden; }
            .header { background-color: {{.PrimaryColor}}; color: #ffffff; padding: 35px 30px 20px 30px; text-align: center; }
            .logo-placeholder { color: {{.SecondaryColor}}; font-size: 20px; font-weight: 900; margin-bottom: 15px; }
            .content { padding: 30px; color: #333333; }
            .score-card { background-color: #f0f4ff; padding: 25px; border-radius: 8px; text-align: center; margin-bottom: 25px; border: 2px solid {{.Color}}; }
            .score-number { font-size: 52px; font-weight: 800; color: {{.Color}}; line-height: 1; }
            .readiness-level { font-size: 22px; font-weight: 700; color: {{.PrimaryColor}}; margin-top: 10px; }
            .rating-badge { display: inline-block; padding: 6px 18px; color: {{.PrimaryColor}}; border-radius: 25px; font-weight: 700; margin-top: 15px; font-size: 14px; text-transform: uppercase; background-color: {{.SecondaryColor}}; }
            .section-title { font-size: 20px; color: {{.PrimaryColor}}; margin-top: 35px; border-bottom: 2px solid {{.SecondaryColor}}; padding-bottom: 8px; font-weight: 700; }
            .insight { background-color: #fff9e6; border-left: 4px solid #ffc107; padding: 15px; margin-top: 20px; border-radius: 4px; font-style: italic; color: #555555; }
            .footer { background-color: #eeeeee; color: #777777; padding: 20px; text-align: center; font-size: 12px; }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <div class="logo-placeholder">AXONVA CONSULTING</div>
                <h1 style="margin:0; font-size:26px;">Your AI Readiness Assessment Results</h1>
            </div>
            <div class="content">
                <p>Hello <strong>{{.UserName}}</strong>,</p>
                <p>Thank you for completing your assessment. We've compiled a summary of your results.</p>
                
                <div class="score-card">
                    <h2 style="margin:0; font-size:18px; color:{{.PrimaryColor}};">Overall AI Readiness Rating:</h2>
                    <p class="score-number">{{.UserScore}}<span style="font-size: 24px; font-weight: 400; color: #333;">/100</span></p>
                    <p class="readiness-level">{{.Emoji}} {{.ReadinessLevel}}</p>
                    <span class="rating-badge">{{.RatingTitle}}</span>
                </div>

                <div class="section-title">Key Insight</div>
                <div class="insight">{{.Insight}}</div>

                <div class="section-title">Tailored Next Steps</div>
                <ul style="list-style-type: none; padding-left: 0; margin-top: 20px;">
                    {{range .ActionSteps}}
                    <li style="margin-bottom: 12px; line-height: 1.5; font-size: 14px; color: #333333; border-left: 3px solid {{$.SecondaryColor}}; padding-left: 10px;">{{.}}</li>
                    {{end}}
                </ul>
                
                <p style="margin-top: 30px;">Best regards,<br><strong>The Axonva Consulting Team</strong></p>
            </div>
            <div class="footer">
                &copy; {{.CurrentYear}} Axonva Consulting.
            </div>
        </div>
    </body>
    </html>
    `

	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

type ContactEmailData struct {
	SenderName     string
	SenderEmail    string
	Message        string
	PrimaryColor   string
	SecondaryColor string
	CurrentYear    int
}

func GenerateContactFormEmail(name, email, message string) (string, error) {
	data := ContactEmailData{
		SenderName:     name,
		SenderEmail:    email,
		Message:        message,
		PrimaryColor:   "#0b3937", // Your brand teal
		SecondaryColor: "#a9fa60", // Your brand lime
		CurrentYear:    time.Now().Year(),
	}

	const tmpl = `
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <style>
            body { font-family: 'Helvetica', Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 20px; }
            .wrapper { max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 8px; overflow: hidden; border: 1px solid #ddd; }
            .header { background-color: {{.PrimaryColor}}; color: white; padding: 30px; text-align: center; }
            .content { padding: 30px; line-height: 1.6; color: #333; }
            .field-label { font-weight: bold; color: {{.PrimaryColor}}; text-transform: uppercase; font-size: 12px; margin-bottom: 5px; }
            .field-value { background: #f9f9f9; padding: 15px; border-radius: 5px; border-left: 4px solid {{.SecondaryColor}}; margin-bottom: 20px; }
            .footer { background: #eeeeee; padding: 20px; text-align: center; font-size: 12px; color: #777; }
            .btn { display: inline-block; padding: 12px 25px; background-color: {{.SecondaryColor}}; color: {{.PrimaryColor}}; text-decoration: none; border-radius: 5px; font-weight: bold; margin-top: 10px; }
        </style>
    </head>
    <body>
        <div class="wrapper">
            <div class="header">
                <div style="font-size: 18px; font-weight: 900; letter-spacing: 2px;">AXONVA</div>
                <h2 style="margin: 10px 0 0 0;">New Inquiry Received</h2>
            </div>
            <div class="content">
                <p>Hello Team, you have a new message from the website contact form.</p>
                
                <div class="field-label">From</div>
                <div class="field-value"><strong>{{.SenderName}}</strong> ({{.SenderEmail}})</div>

                <div class="field-label">Message Details</div>
                <div class="field-value" style="white-space: pre-wrap;">{{.Message}}</div>

                <div style="text-align: center;">
                    <a href="mailto:{{.SenderEmail}}" class="btn">Reply to {{.SenderName}}</a>
                </div>
            </div>
            <div class="footer">
                &copy; {{.CurrentYear}} Axonva Consulting Internal Notification.
            </div>
        </div>
    </body>
    </html>
    `

	t, err := template.New("contactEmail").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type ServiceRequestData struct {
	Name             string
	Email            string
	Phone            string
	ServiceType      string
	RequestedModules []string
	PreferredDate    string
	Message          string
	PrimaryColor     string
	SecondaryColor   string
	CurrentYear      int
}

func GenerateServiceRequestEmail(data ServiceRequestData) (string, error) {
	// Default Branding
	data.PrimaryColor = "#0b3937"
	data.SecondaryColor = "#a9fa60"
	data.CurrentYear = time.Now().Year()

	const tmpl = `
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <style>
            body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f8f9fa; margin: 0; padding: 0; }
            .container { max-width: 600px; margin: 20px auto; background: #ffffff; border-radius: 12px; overflow: hidden; border: 1px solid #e0e0e0; }
            .header { background-color: {{.PrimaryColor}}; color: #ffffff; padding: 40px 20px; text-align: center; }
            .status-badge { display: inline-block; background-color: {{.SecondaryColor}}; color: {{.PrimaryColor}}; padding: 5px 15px; border-radius: 20px; font-size: 12px; font-weight: bold; text-transform: uppercase; margin-bottom: 10px; }
            .content { padding: 30px; color: #333333; }
            .info-grid { width: 100%; border-collapse: collapse; margin-bottom: 20px; }
            .info-row td { padding: 10px; border-bottom: 1px solid #f0f0f0; vertical-align: top; }
            .label { font-weight: bold; color: {{.PrimaryColor}}; width: 140px; font-size: 13px; text-transform: uppercase; }
            .value { color: #555555; font-size: 15px; }
            .module-tag { display: inline-block; background: #f0f4ff; color: {{.PrimaryColor}}; padding: 2px 8px; border-radius: 4px; margin: 2px; font-size: 12px; border: 1px solid #d1d9e6; }
            .message-box { background-color: #f9f9f9; padding: 20px; border-radius: 8px; border-left: 4px solid {{.PrimaryColor}}; font-style: italic; margin-top: 10px; }
            .footer { background-color: #f4f4f4; color: #888888; padding: 20px; text-align: center; font-size: 12px; }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <span class="status-badge">New Service Request</span>
                <h1 style="margin: 0; font-size: 24px;">{{.ServiceType}}</h1>
            </div>
            <div class="content">
                <table class="info-grid">
                    <tr class="info-row">
                        <td class="label">Client Name</td>
                        <td class="value"><strong>{{.Name}}</strong></td>
                    </tr>
                    <tr class="info-row">
                        <td class="label">Email</td>
                        <td class="value">{{.Email}}</td>
                    </tr>
                    <tr class="info-row">
                        <td class="label">Phone</td>
                        <td class="value">{{.Phone}}</td>
                    </tr>
                    <tr class="info-row">
                        <td class="label">Meeting Date</td>
                        <td class="value" style="color: #d9534f; font-weight: bold;">{{.PreferredDate}}</td>
                    </tr>
                    {{if .RequestedModules}}
                    <tr class="info-row">
                        <td class="label">Modules</td>
                        <td class="value">
                            {{range .RequestedModules}}
                                <span class="module-tag">{{.}}</span>
                            {{end}}
                        </td>
                    </tr>
                    {{end}}
                </table>

                <div class="label" style="margin-top: 25px;">Client Message:</div>
                <div class="message-box">
                    "{{.Message}}"
                </div>

                <div style="margin-top: 30px; text-align: center;">
                    <a href="mailto:{{.Email}}" style="background-color: {{.PrimaryColor}}; color: white; padding: 12px 25px; text-decoration: none; border-radius: 6px; font-weight: bold; font-size: 14px;">Review & Respond</a>
                </div>
            </div>
            <div class="footer">
                This request was generated via the Axonva Consulting Service Portal.
            </div>
        </div>
    </body>
    </html>
    `

	t, err := template.New("serviceRequest").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type SenderReceiptData struct {
	Name           string
	FormType       string // "Inquiry" or "Service Request"
	PrimaryColor   string
	SecondaryColor string
	CurrentYear    int
}

func GenerateSenderAcknowledgment(name string, isServiceRequest bool) (string, error) {
	formType := "Inquiry"
	if isServiceRequest {
		formType = "Service Request"
	}

	data := SenderReceiptData{
		Name:           name,
		FormType:       formType,
		PrimaryColor:   "#0b3937",
		SecondaryColor: "#a9fa60",
		CurrentYear:    time.Now().Year(),
	}

	const tmpl = `
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <style>
            body { font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; background-color: #f9f9f9; margin: 0; padding: 0; }
            .email-container { max-width: 600px; margin: 40px auto; background: #ffffff; border-radius: 8px; overflow: hidden; border: 1px solid #eeeeee; }
            .banner { background-color: {{.PrimaryColor}}; padding: 30px; text-align: center; }
            .logo { color: {{.SecondaryColor}}; font-size: 24px; font-weight: 900; letter-spacing: 2px; }
            .body-content { padding: 40px; color: #333333; line-height: 1.7; }
            .highlight-box { border-left: 4px solid {{.SecondaryColor}}; background-color: #f0f7f7; padding: 20px; margin: 25px 0; border-radius: 0 8px 8px 0; }
            .footer { background-color: #f4f4f4; padding: 20px; text-align: center; font-size: 12px; color: #999999; }
            .btn { display: inline-block; margin-top: 20px; padding: 12px 24px; background-color: {{.PrimaryColor}}; color: #ffffff; text-decoration: none; border-radius: 4px; font-weight: bold; }
        </style>
    </head>
    <body>
        <div class="email-container">
            <div class="banner">
                <div class="logo">AXONVA</div>
            </div>
            <div class="body-content">
                <h2 style="margin-top: 0; color: {{.PrimaryColor}};">We’ve received your {{.FormType}}</h2>
                <p>Hello <strong>{{.Name}}</strong>,</p>
                
                <p>Thank you for reaching out to Axonva Consulting. We wanted to let you know that <strong>we have received your email</strong> and our team is currently reviewing the details.</p>
                
                <div class="highlight-box">
                    <p style="margin: 0; font-weight: 500;">
                        We will get back to you as soon as possible, typically within one business day, to discuss the next steps.
                    </p>
                </div>

                <p>If your matter is urgent, please feel free to reply to this thread or give us a call.</p>
                
                <p style="margin-top: 40px;">Best regards,<br>
                <strong>Axonva Consulting Team</strong></p>
            </div>
            <div class="footer">
                &copy; {{.CurrentYear}} Axonva Consulting. AI Consultation Done Right.
            </div>
        </div>
    </body>
    </html>
    `

	t, err := template.New("senderAck").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
