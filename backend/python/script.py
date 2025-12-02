import spacy
import re
import dateparser

nlp = spacy.load("en_core_web_sm")

email_text = """Dear Mr. Thompson,

I hope this message finds you well. My name is Caroline Whitaker, and I am the Marketing Director at Northbridge Consulting Group. I am writing to formally request your services for the creation of a YouTube video for our company’s upcoming digital outreach initiative.

Below is an overview of the project requirements:

Topic/Goal: An informational video introducing our new business analytics platform, Northbridge Insights, aimed at educating potential clients on key features and benefits.

Length: Approximately 4–6 minutes.

Style: Professional and promotional, incorporating a clean, modern aesthetic with light motion graphics.

Deadline: Preferably by March 18, 2026.

Additional Details: We would require both scripting and voiceover services. Our team will provide brand guidelines, logos, sample footage, and any necessary documentation.

We are willing to pay you 1000$ for this. We would be grateful for the opportunity to collaborate with you and look forward to hearing your recommendations.

Thank you for your time and consideration.

Sincerely,
Caroline Whitaker
Marketing Director, Northbridge Consulting Group
caroline.whitaker@northbridgecg.com"""

def extract_email_data(email_text):
    data = {
        "name": None,
        "email": None,
        "organization": None,
        "location": None,
        "project_details": None,
        "budget": None,
        "deadline": None,
    }

    # ---Extract Name---
    doc = nlp(email_text)
    data["name"] = [ent.text for ent in doc.ents if ent.label_ == "PERSON"]
    
    # ---Extract Location---
    data["location"] = [ent.text for ent in doc.ents if ent.label_ in ["GPE", "LOC"]]
    
    # ---Extract Organization---
    data["organization"] = [ent.text for ent in doc.ents if ent.label_ in ["ORG"]]

    # --- 2. Extract Email Address ---
    email_match = re.search(r'[\w\.-]+@[\w\.-]+', email_text)
    if email_match:
        data["email"] = email_match.group(0)

    # --- 3. Extract Money / Budget ---
    data["budget"] = [ent.text for ent in doc.ents if ent.label_ == "MONEY"]

    # --- 4. Extract Dates / Deadlines ---
    data["deadline"] = [ent.text for ent in doc.ents if ent.label_ == "DATE"]

    # --- 5. Extract Project Details (simple heuristic: sentences with keywords) ---
    data["project_details"] = [ent.text for ent in doc.ents if ent.label_ == "PRODUCT"]

    return data
    
extracted_data = extract_email_data(email_text)
print(extracted_data)