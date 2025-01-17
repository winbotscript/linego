U
    �^�^��  �                   @   s   d dl mZmZmZmZmZ d dlmZ d dlm	Z	 d dl
Z
d dlZddlT d dl mZ d dlmZ g ZG d	d
� d
e�ZG dd� de�ZG dd� dee�ZG dd� de�Ze�e� ddejdedgdffe_G dd� de�Ze�e� d ejdedgdfdejdedgdffe_G dd� de�Ze�e� ddejdedgdffe_G dd� de�Ze�e� d ejde dgdfdejdedgdffe_G dd� de�Z!e�e!� ddejde"dgdffe!_G dd� de�Z#e�e#� d ejde$dgdfdejdedgdffe#_G dd� de�Z%e�e%� ddejde&dgdffe%_G d d!� d!e�Z'e�e'� d ejde(dgdfdejdedgdffe'_G d"d#� d#e�Z)e�e)� ddejde*dgdffe)_G d$d%� d%e�Z+e�e+� d ejde,dgdfdejdedgdffe+_e	e� [dS )&�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolException)�fix_specN�   )�*)�
TProcessor)�
TTransportc                   @   s4   e Zd Zdd� Zdd� Zdd� Zdd� Zd	d
� ZdS )�Ifacec                 C   s   dS ��1
        Parameters:
         - request

        N� ��self�requestr   r   �?/storage/sdcard0/bot1/flex/curve/SecondaryQrCodeLoginService.py�createSession   s    zIface.createSessionc                 C   s   dS r   r   r   r   r   r   �createQrCode   s    zIface.createQrCodec                 C   s   dS r   r   r   r   r   r   �verifyCertificate&   s    zIface.verifyCertificatec                 C   s   dS r   r   r   r   r   r   �createPinCode.   s    zIface.createPinCodec                 C   s   dS r   r   r   r   r   r   �qrCodeLogin6   s    zIface.qrCodeLoginN)�__name__�
__module__�__qualname__r   r   r   r   r   r   r   r   r   r      s
   r   c                   @   s�   e Zd Zd"dd�Zdd� Zdd� Zdd	� Zd
d� Zdd� Zdd� Z	dd� Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� Zdd� Zd d!� ZdS )#�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   �iprot�oprotr   r   r   �__init__@   s    zClient.__init__c                 C   s   | � |� | �� S �r   )�send_createSession�recv_createSessionr   r   r   r   r   F   s    
zClient.createSessionc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   �writeMessageBeginr   �CALLr    �createSession_argsr   �write�writeMessageEnd�trans�flush�r   r   �argsr   r   r   r%   O   s    
zClient.send_createSessionc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz$createSession failed: unknown result)r   �readMessageBeginr   �	EXCEPTIONr   �read�readMessageEnd�createSession_result�success�e�MISSING_RESULT�r   r!   �fname�mtype�rseqid�x�resultr   r   r   r&   W   s    




zClient.recv_createSessionc                 C   s   | � |� | �� S r$   )�send_createQrCode�recv_createQrCoder   r   r   r   r   h   s    
zClient.createQrCodec                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   r'   r   r(   r    �createQrCode_argsr   r*   r+   r,   r-   r.   r   r   r   r>   q   s    
zClient.send_createQrCodec                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz#createQrCode failed: unknown result)r   r0   r   r1   r   r2   r3   �createQrCode_resultr5   r6   r7   r8   r   r   r   r?   y   s    




zClient.recv_createQrCodec                 C   s   | � |� | �� S r$   )�send_verifyCertificate�recv_verifyCertificater   r   r   r   r   �   s    
zClient.verifyCertificatec                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   r'   r   r(   r    �verifyCertificate_argsr   r*   r+   r,   r-   r.   r   r   r   rB   �   s    
zClient.send_verifyCertificatec                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz(verifyCertificate failed: unknown result)r   r0   r   r1   r   r2   r3   �verifyCertificate_resultr5   r6   r7   r8   r   r   r   rC   �   s    




zClient.recv_verifyCertificatec                 C   s   | � |� | �� S r$   )�send_createPinCode�recv_createPinCoder   r   r   r   r   �   s    
zClient.createPinCodec                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   r'   r   r(   r    �createPinCode_argsr   r*   r+   r,   r-   r.   r   r   r   rF   �   s    
zClient.send_createPinCodec                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz$createPinCode failed: unknown result)r   r0   r   r1   r   r2   r3   �createPinCode_resultr5   r6   r7   r8   r   r   r   rG   �   s    




zClient.recv_createPinCodec                 C   s   | � |� | �� S r$   )�send_qrCodeLogin�recv_qrCodeLoginr   r   r   r   r   �   s    
zClient.qrCodeLoginc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   r'   r   r(   r    �qrCodeLogin_argsr   r*   r+   r,   r-   r.   r   r   r   rJ   �   s    
zClient.send_qrCodeLoginc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz"qrCodeLogin failed: unknown result)r   r0   r   r1   r   r2   r3   �qrCodeLogin_resultr5   r6   r7   r8   r   r   r   rK   �   s    




zClient.recv_qrCodeLogin)N)r   r   r   r#   r   r%   r&   r   r>   r?   r   rB   rC   r   rF   rG   r   rJ   rK   r   r   r   r   r   ?   s    
					r   c                   @   sL   e Zd Zdd� Zdd� Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )�	Processorc                 C   sR   || _ i | _tj| jd< tj| jd< tj| jd< tj| jd< tj| jd< d | _d S )Nr   r   r   r   r   )	�_handler�_processMaprN   �process_createSession�process_createQrCode�process_verifyCertificate�process_createPinCode�process_qrCodeLogin�_on_message_begin)r   �handlerr   r   r   r#   �   s    zProcessor.__init__c                 C   s
   || _ d S �N)rV   )r   �funcr   r   r   �on_message_begin�   s    zProcessor.on_message_beginc                 C   s�   |� � \}}}| jr"| �|||� || jkr�|�tj� |��  ttjd| �}|�	|t
j|� |�|� |��  |j��  d S | j| | |||� dS )NzUnknown function %sT)r0   rV   rP   �skipr   �STRUCTr3   r   �UNKNOWN_METHODr'   r   r1   r*   r+   r,   r-   )r   r!   r"   �name�type�seqidr<   r   r   r   �process�   s    


zProcessor.processc           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )N�!TApplication exception in handler�Unexpected exception in handler�Internal errorr   )r)   r2   r3   r4   rO   r   r   r5   r   �REPLYr   �TTransportException�SecondaryQrCodeExceptionr6   r   �logging�	exceptionr1   �	Exception�INTERNAL_ERRORr'   r*   r+   r,   r-   �	r   r`   r!   r"   r/   r=   �msg_typer6   Zexr   r   r   rQ     s0    




zProcessor.process_createSessionc           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )Nrb   rc   rd   r   )r@   r2   r3   rA   rO   r   r   r5   r   re   r   rf   rg   r6   r   rh   ri   r1   rj   rk   r'   r*   r+   r,   r-   rl   r   r   r   rR   *  s0    




zProcessor.process_createQrCodec           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )Nrb   rc   rd   r   )rD   r2   r3   rE   rO   r   r   r5   r   re   r   rf   rg   r6   r   rh   ri   r1   rj   rk   r'   r*   r+   r,   r-   rl   r   r   r   rS   D  s0    




z#Processor.process_verifyCertificatec           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )Nrb   rc   rd   r   )rH   r2   r3   rI   rO   r   r   r5   r   re   r   rf   rg   r6   r   rh   ri   r1   rj   rk   r'   r*   r+   r,   r-   rl   r   r   r   rT   ^  s0    




zProcessor.process_createPinCodec           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )Nrb   rc   rd   r   )rL   r2   r3   rM   rO   r   r   r5   r   re   r   rf   rg   r6   r   rh   ri   r1   rj   rk   r'   r*   r+   r,   r-   rl   r   r   r   rU   x  s0    




zProcessor.process_qrCodeLoginN)r   r   r   r#   rZ   ra   rQ   rR   rS   rT   rU   r   r   r   r   rN   �   s   
rN   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r)   �%
    Attributes:
     - request

    Nc                 C   s
   || _ d S rX   �r   r   r   r   r   r#   �  s    zcreateSession_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr	   )�_fast_decode�
isinstancer,   r   �CReadableTransport�thrift_spec�	__class__�readStructBegin�readFieldBeginr   �STOPr\   �"LoginQrCode_CreateQrSessionRequestr   r2   r[   �readFieldEnd�readStructEnd�r   r!   r9   �ftype�fidr   r   r   r2   �  s    "



zcreateSession_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr)   r   r	   ��_fast_encodert   r,   r*   ru   �writeStructBeginr   �writeFieldBeginr   r\   �writeFieldEnd�writeFieldStop�writeStructEnd�r   r"   r   r   r   r*   �  s    

zcreateSession_args.writec                 C   s   d S rX   r   �r   r   r   r   �validate�  s    zcreateSession_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS �z%s=%rr   ��.0�key�valuer   r   r   �
<listcomp>�  s   �z/createSession_args.__repr__.<locals>.<listcomp>�%s(%s)�, ��__dict__�itemsru   r   �join�r   �Lr   r   r   �__repr__�  s    �zcreateSession_args.__repr__c                 C   s   t || j�o| j|jkS rX   �rr   ru   r�   �r   �otherr   r   r   �__eq__�  s    zcreateSession_args.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   �__ne__�  s    zcreateSession_args.__ne__)N�r   r   r   �__doc__r#   r2   r*   r�   r�   r�   r�   r   r   r   r   r)   �  s   
r)   r   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r4   �.
    Attributes:
     - success
     - e

    Nc                 C   s   || _ || _d S rX   �r5   r6   �r   r5   r6   r   r   r   r#   �  s    zcreateSession_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr   r	   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �#LoginQrCode_CreateQrSessionResponser5   r2   r[   rg   r6   rz   r{   r|   r   r   r   r2   �  s(    "




zcreateSession_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr4   r5   r   r6   r	   �r�   rt   r,   r*   ru   r�   r5   r�   r   r\   r�   r6   r�   r�   r�   r   r   r   r*   �  s    


zcreateSession_result.writec                 C   s   d S rX   r   r�   r   r   r   r�     s    zcreateSession_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�     s   �z1createSession_result.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�     s    �zcreateSession_result.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�     s    zcreateSession_result.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�     s    zcreateSession_result.__ne__)NNr�   r   r   r   r   r4   �  s   
r4   r5   r6   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r@   rn   Nc                 C   s
   || _ d S rX   ro   r   r   r   r   r#   '  s    zcreateQrCode_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S rp   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �LoginQrCode_CreateQrCodeRequestr   r2   r[   rz   r{   r|   r   r   r   r2   *  s    "



zcreateQrCode_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr@   r   r	   r   r�   r   r   r   r*   >  s    

zcreateQrCode_args.writec                 C   s   d S rX   r   r�   r   r   r   r�   J  s    zcreateQrCode_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   N  s   �z.createQrCode_args.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   M  s    �zcreateQrCode_args.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   R  s    zcreateQrCode_args.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   U  s    zcreateQrCode_args.__ne__)Nr�   r   r   r   r   r@     s   
r@   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rA   r�   Nc                 C   s   || _ || _d S rX   r�   r�   r   r   r   r#   g  s    zcreateQrCode_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   � LoginQrCode_CreateQrCodeResponser5   r2   r[   rg   r6   rz   r{   r|   r   r   r   r2   k  s(    "




zcreateQrCode_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrA   r5   r   r6   r	   r�   r�   r   r   r   r*   �  s    


zcreateQrCode_result.writec                 C   s   d S rX   r   r�   r   r   r   r�   �  s    zcreateQrCode_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   �  s   �z0createQrCode_result.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   �  s    �zcreateQrCode_result.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   �  s    zcreateQrCode_result.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   �  s    zcreateQrCode_result.__ne__)NNr�   r   r   r   r   rA   ^  s   
rA   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rD   rn   Nc                 C   s
   || _ d S rX   ro   r   r   r   r   r#   �  s    zverifyCertificate_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S rp   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �$LoginQrCode_VerifyCertificateRequestr   r2   r[   rz   r{   r|   r   r   r   r2   �  s    "



zverifyCertificate_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )NrD   r   r	   r   r�   r   r   r   r*   �  s    

zverifyCertificate_args.writec                 C   s   d S rX   r   r�   r   r   r   r�   �  s    zverifyCertificate_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   �  s   �z3verifyCertificate_args.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   �  s    �zverifyCertificate_args.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   �  s    zverifyCertificate_args.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   �  s    zverifyCertificate_args.__ne__)Nr�   r   r   r   r   rD   �  s   
rD   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rE   r�   Nc                 C   s   || _ || _d S rX   r�   r�   r   r   r   r#   �  s    z!verifyCertificate_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �%LoginQrCode_VerifyCertificateResponser5   r2   r[   rg   r6   rz   r{   r|   r   r   r   r2   �  s(    "




zverifyCertificate_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrE   r5   r   r6   r	   r�   r�   r   r   r   r*     s    


zverifyCertificate_result.writec                 C   s   d S rX   r   r�   r   r   r   r�     s    z!verifyCertificate_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   #  s   �z5verifyCertificate_result.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   "  s    �z!verifyCertificate_result.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   '  s    zverifyCertificate_result.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   *  s    zverifyCertificate_result.__ne__)NNr�   r   r   r   r   rE   �  s   
rE   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rH   rn   Nc                 C   s
   || _ d S rX   ro   r   r   r   r   r#   ;  s    zcreatePinCode_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S rp   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   � LoginQrCode_CreatePinCodeRequestr   r2   r[   rz   r{   r|   r   r   r   r2   >  s    "



zcreatePinCode_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )NrH   r   r	   r   r�   r   r   r   r*   R  s    

zcreatePinCode_args.writec                 C   s   d S rX   r   r�   r   r   r   r�   ^  s    zcreatePinCode_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   b  s   �z/createPinCode_args.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   a  s    �zcreatePinCode_args.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   f  s    zcreatePinCode_args.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   i  s    zcreatePinCode_args.__ne__)Nr�   r   r   r   r   rH   3  s   
rH   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rI   r�   Nc                 C   s   || _ || _d S rX   r�   r�   r   r   r   r#   {  s    zcreatePinCode_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �!LoginQrCode_CreatePinCodeResponser5   r2   r[   rg   r6   rz   r{   r|   r   r   r   r2     s(    "




zcreatePinCode_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrI   r5   r   r6   r	   r�   r�   r   r   r   r*   �  s    


zcreatePinCode_result.writec                 C   s   d S rX   r   r�   r   r   r   r�   �  s    zcreatePinCode_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   �  s   �z1createPinCode_result.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   �  s    �zcreatePinCode_result.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   �  s    zcreatePinCode_result.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   �  s    zcreatePinCode_result.__ne__)NNr�   r   r   r   r   rI   r  s   
rI   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rL   rn   Nc                 C   s
   || _ d S rX   ro   r   r   r   r   r#   �  s    zqrCodeLogin_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S rp   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �LoginQrCode_QrCodeLoginRequestr   r2   r[   rz   r{   r|   r   r   r   r2   �  s    "



zqrCodeLogin_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )NrL   r   r	   r   r�   r   r   r   r*   �  s    

zqrCodeLogin_args.writec                 C   s   d S rX   r   r�   r   r   r   r�   �  s    zqrCodeLogin_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   �  s   �z-qrCodeLogin_args.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   �  s    �zqrCodeLogin_args.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   �  s    zqrCodeLogin_args.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   �  s    zqrCodeLogin_args.__ne__)Nr�   r   r   r   r   rL   �  s   
rL   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )rM   r�   Nc                 C   s   || _ || _d S rX   r�   r�   r   r   r   r#     s    zqrCodeLogin_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   )rq   rr   r,   r   rs   rt   ru   rv   rw   r   rx   r\   �LoginQrCode_QrCodeLoginResponser5   r2   r[   rg   r6   rz   r{   r|   r   r   r   r2   	  s(    "




zqrCodeLogin_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrM   r5   r   r6   r	   r�   r�   r   r   r   r*   #  s    


zqrCodeLogin_result.writec                 C   s   d S rX   r   r�   r   r   r   r�   3  s    zqrCodeLogin_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r�   r   r�   r   r   r   r�   7  s   �z/qrCodeLogin_result.__repr__.<locals>.<listcomp>r�   r�   r�   r�   r   r   r   r�   6  s    �zqrCodeLogin_result.__repr__c                 C   s   t || j�o| j|jkS rX   r�   r�   r   r   r   r�   ;  s    zqrCodeLogin_result.__eq__c                 C   s
   | |k S rX   r   r�   r   r   r   r�   >  s    zqrCodeLogin_result.__ne__)NNr�   r   r   r   r   rM   �  s   
rM   )-�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   �thrift.TRecursiver   �sysrh   �ttypesr   �thrift.transportr   �all_structs�objectr   r   rN   r)   �appendr\   ry   rt   r4   r�   rg   r@   r�   rA   r�   rD   r�   rE   r�   rH   r�   rI   r�   rL   r�   rM   r�   r   r   r   r   �<module>	   s�   * 3 %8
�D
�8
�D
�8
�D
�8
�D
�8
�D
�