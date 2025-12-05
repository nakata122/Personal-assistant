import spacy
import re
import dateparser
import sys

nlp = spacy.load("en_core_web_lg")

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
    data["name"] = {ent.text for ent in doc.ents if ent.label_ == "PERSON"}
    
    # ---Extract Location---
    data["location"] = {ent.text for ent in doc.ents if ent.label_ in ["GPE", "LOC"]}
    
    # ---Extract Organization---
    data["organization"] = {ent.text for ent in doc.ents if ent.label_ in ["ORG"]}

    # --- 2. Extract Email Address ---
    email_match = re.search(r'[\w\.-]+@[\w\.-]+', email_text)
    if email_match:
        data["email"] = email_match.group(0)

    # --- 3. Extract Money / Budget ---
    data["budget"] = {ent.text for ent in doc.ents if ent.label_ == "MONEY"}

    # --- 4. Extract Dates / Deadlines ---
    data["deadline"] = {ent.text for ent in doc.ents if ent.label_ == "DATE"}

    # --- 5. Extract Project Details (simple heuristic: sentences with keywords) ---
    data["project_details"] = {ent.text for ent in doc.ents if ent.label_ == "PRODUCT"}

    return data
    
print(sys.argv[1])
extracted_data = extract_email_data(sys.argv[1])
print(extracted_data)