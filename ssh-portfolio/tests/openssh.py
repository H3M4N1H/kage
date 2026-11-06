import subprocess,pathlib
key=str(pathlib.Path(__file__).resolve().parents[1] / '.state/host_ed25519')
pub=subprocess.check_output(['ssh-keygen','-y','-f',key],text=True).strip()
known=pathlib.Path('/tmp/portfolio-known-hosts');known.write_text('[127.0.0.1]:2222 '+pub+'\n')
base=['ssh','-p','2222','-o','StrictHostKeyChecking=yes','-o',f'UserKnownHostsFile={known}','-o','BatchMode=yes','-o','ConnectTimeout=3']
p=subprocess.run(base+['-T','visitor@127.0.0.1'],capture_output=True,text=True,timeout=8)
assert p.returncode==0,p.stderr
assert 'CTF-AI' in p.stdout and 'Hemanth' in p.stdout
p=subprocess.run(base+['visitor@127.0.0.1','cat /etc/passwd'],capture_output=True,text=True,timeout=8)
assert p.returncode!=0 and 'root:' not in p.stdout
p=subprocess.run(base+['-s','visitor@127.0.0.1','sftp'],capture_output=True,text=True,timeout=8)
assert p.returncode!=0
print('OpenSSH: pinned host key verified; plain portfolio passed; exec and SFTP rejected.')
