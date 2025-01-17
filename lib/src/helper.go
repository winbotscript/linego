U
    rܢ^�<  �                   @   s�   d dl mZmZmZmZmZ d dlmZ d dlZd dl	Z	ddl
mZmZmZ d dl mZ d dlmZ G dd	� d	e�ZG d
d� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZdS )�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolExceptionN�   )�	GroupCall�GroupCallRoute�TalkException)�
TProcessor)�
TTransportc                   @   s@   e Zd ZdZdddejdddffZddd�Zdd	� Zd
d� Z	dS )�getGroupCall_argsz$
    Attributes:
     - chatMid
    N�   �chatMid�UTF8c                 C   s
   || _ d S �N)r   )�selfr   � r   �//storage/sdcard0/bot1/flex/curve/CallService.py�__init__   s    zgetGroupCall_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� n
|�|� |��  q6|��  d S )Nr   r   �utf-8)�_fast_decode�thrift_spec�	__class__�readStructBegin�readFieldBeginr   �STOP�STRING�sys�version_info�
readString�decoder   �skip�readFieldEnd�readStructEnd�r   �iprot�fname�ftype�fidr   r   r   �read   s    

(

zgetGroupCall_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )Nr   r   r   r   r   )�_fast_encoder   �trans�writer   �writeStructBeginr   �writeFieldBeginr   r   �writeStringr   r    �encode�writeFieldEnd�writeFieldStop�writeStructEnd�r   �oprotr   r   r   r.   -   s    

&zgetGroupCall_args.write)N)
�__name__�
__module__�__qualname__�__doc__r   r   r   r   r+   r.   r   r   r   r   r      s   �
r   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�getGroupCall_result�-
    Attributes:
     - success
     - e
    r   �successNr   �ec                 C   s   || _ || _d S r   �r>   r?   �r   r>   r?   r   r   r   r   E   s    zgetGroupCall_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S �Nr   r   )r   r   r   r   r   r   r   �STRUCTr	   r>   r+   r#   r   r?   r$   r%   r&   r   r   r   r+   I   s(    




zgetGroupCall_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr<   r>   r   r?   r   �r,   r   r-   r.   r   r/   r>   r0   r   rC   r3   r?   r4   r5   r6   r   r   r   r.   c   s    


zgetGroupCall_result.write)NN)r8   r9   r:   r;   r   rC   r	   r   r   r   r+   r.   r   r   r   r   r<   :   s   �
r<   c                	   @   sd   e Zd ZdZdddejdddfdejdejddfdfd	ejd
ddffZddd�Z	dd� Z
dd� ZdS )�inviteIntoGroupCall_argszG
    Attributes:
     - chatMid
     - memberMids
     - mediaType
    Nr   r   r   �   �
memberMidsF�   �	mediaTypec                 C   s   || _ || _|| _d S r   )r   rG   rI   )r   r   rG   rI   r   r   r   r   �   s    z!inviteIntoGroupCall_args.__init__c           	      C   sb  |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrR�qV|dkr�|tjkr�tj	d dkr�|�
� �d�n|�
� | _n
|�|� n�|dk�r|tjk�rg | _|�� \}}t|�D ]4}tj	d dkr�|�
� �d�n|�
� }| j�|� q�|��  n
|�|� n8|dk�rB|tjk�r6|�� | _n
|�|� n
|�|� |��  q6|��  d S )Nr   r   r   rF   rH   )r   r   r   r   r   r   r   r   r   r    r!   r"   r   r#   �LISTrG   �readListBegin�range�append�readListEnd�I32�readI32rI   r$   r%   )	r   r'   r(   r)   r*   Z
_etype2189Z	_size2186Z_i2190Z	_elem2191r   r   r   r+   �   s6    

(
$



zinviteIntoGroupCall_args.readc                 C   s6  |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� |�tj	t| j�� | jD ]&}|�
tjd dkr�|�d�n|� q�|��  |��  | jd k	�r"|�dtjd	� |�| j� |��  |��  |��  d S )
NrE   r   r   r   r   rG   rF   rI   rH   )r,   r   r-   r.   r   r/   r   r0   r   r   r1   r   r    r2   r3   rG   rJ   �writeListBegin�len�writeListEndrI   rO   �writeI32r4   r5   )r   r7   Ziter2192r   r   r   r.   �   s*    

&

$zinviteIntoGroupCall_args.write)NNN)r8   r9   r:   r;   r   r   rJ   rO   r   r   r+   r.   r   r   r   r   rE   t   s   �
"rE   c                   @   sB   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dS )�inviteIntoGroupCall_resultz
    Attributes:
     - e
    Nr   r?   c                 C   s
   || _ d S r   )r?   )r   r?   r   r   r   r   �   s    z#inviteIntoGroupCall_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S )Nr   )r   r   r   r   r   r   r   rC   r   r?   r+   r#   r$   r%   r&   r   r   r   r+   �   s    



zinviteIntoGroupCall_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )NrU   r?   r   )r,   r   r-   r.   r   r/   r?   r0   r   rC   r3   r4   r5   r6   r   r   r   r.   �   s    

z inviteIntoGroupCall_result.write)N)r8   r9   r:   r;   r   rC   r   r   r   r+   r.   r   r   r   r   rU   �   s   �
rU   c                   @   sN   e Zd ZdZdddejdddfdejdddffZddd	�Zd
d� Z	dd� Z
dS )�acquireGroupCallRoute_argsz5
    Attributes:
     - chatMid
     - mediaType
    Nr   r   r   rF   rI   c                 C   s   || _ || _d S r   )r   rI   )r   r   rI   r   r   r   r   �   s    z#acquireGroupCallRoute_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� n4|dkr�|tjkr�|�� | _q�|�|� n
|�|� |��  q6|��  d S )Nr   r   r   rF   )r   r   r   r   r   r   r   r   r   r    r!   r"   r   r#   rO   rP   rI   r$   r%   r&   r   r   r   r+     s$    

(


zacquireGroupCallRoute_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtjd� |�| j� |��  |��  |��  d S )NrV   r   r   r   r   rI   rF   )r,   r   r-   r.   r   r/   r   r0   r   r   r1   r   r    r2   r3   rI   rO   rT   r4   r5   r6   r   r   r   r.     s    

&
z acquireGroupCallRoute_args.write)NN)r8   r9   r:   r;   r   r   rO   r   r   r+   r.   r   r   r   r   rV   �   s   �
rV   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�acquireGroupCallRoute_resultr=   r   r>   Nr   r?   c                 C   s   || _ || _d S r   r@   rA   r   r   r   r   5  s    z%acquireGroupCallRoute_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S rB   )r   r   r   r   r   r   r   rC   r
   r>   r+   r#   r   r?   r$   r%   r&   r   r   r   r+   9  s(    




z!acquireGroupCallRoute_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrW   r>   r   r?   r   rD   r6   r   r   r   r.   S  s    


z"acquireGroupCallRoute_result.write)NN)r8   r9   r:   r;   r   rC   r
   r   r   r   r+   r.   r   r   r   r   rW   *  s   �
rW   c                   @   s.   e Zd Zd
dd�Zdd� Zdd� Zdd	� ZdS )�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   r'   r7   r   r   r   r   d  s    zClient.__init__c           	      C   s�   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  | j}|�� \}}}|tjkr|t� }|�|� |��  |�t� }|�|� |��  |jdk	r�|jS |jdk	r�|j�ttjd��dS )z0
        Parameters:
         - chatMid
        �getGroupCallNz#getGroupCall failed: unknown result)rZ   �writeMessageBeginr   �CALLr[   r   r   r.   �writeMessageEndr-   �flushrY   �readMessageBegin�	EXCEPTIONr   r+   �readMessageEndr<   r>   r?   �MISSING_RESULT)	r   r   �argsr'   r(   �mtype�rseqid�x�resultr   r   r   r\   j  s*    





zClient.getGroupCallc                 C   s�   | j �dtj| j� t� }||_||_||_|�	| j � | j �
�  | j j��  | j}|�� \}}}|tjkr�t� }	|	�|� |��  |	�t� }
|
�|� |��  |
jdk	r�|
j�dS )z[
        Parameters:
         - chatMid
         - memberMids
         - mediaType
        �inviteIntoGroupCallN)rZ   r]   r   r^   r[   rE   r   rG   rI   r.   r_   r-   r`   rY   ra   rb   r   r+   rc   rU   r?   )r   r   rG   rI   re   r'   r(   rf   rg   rh   ri   r   r   r   rj   �  s*    




zClient.inviteIntoGroupCallc           
      C   s�   | j �dtj| j� t� }||_||_|�| j � | j �	�  | j j
��  | j}|�� \}}}|tjkr�t� }|�|� |��  |�t� }	|	�|� |��  |	jdk	r�|	jS |	jdk	r�|	j�ttjd��dS )zE
        Parameters:
         - chatMid
         - mediaType
        �acquireGroupCallRouteNz,acquireGroupCallRoute failed: unknown result)rZ   r]   r   r^   r[   rV   r   rI   r.   r_   r-   r`   rY   ra   rb   r   r+   rc   rW   r>   r?   rd   )
r   r   rI   re   r'   r(   rf   rg   rh   ri   r   r   r   rk   �  s,    





zClient.acquireGroupCallRoute)N)r8   r9   r:   r   r\   rj   rk   r   r   r   r   rX   c  s   
rX   )�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   r   �logging�ttypesr	   r
   r   r   �thrift.transportr   �objectr   r<   rE   rU   rV   rW   rX   r   r   r   r   �<module>   s   .:N.:9