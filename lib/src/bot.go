U
    tܢ^�Z  �                   @   s  d dl mZmZmZmZmZ d dlmZ d dlZd dl	Z	ddl
mZmZmZmZ d dl mZ d dlmZ G dd	� d	e�ZG d
d� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZG dd� de�ZdS )�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolExceptionN�   )�ChannelInfo�ChannelInfos�ChannelToken�ChannelException)�
TProcessor)�
TTransportc                   @   s>   e Zd ZdZddejdddffZddd�Zdd	� Zd
d� Z	dS )�issueChannelToken_args�&
    Attributes:
     - channelId
    Nr   �	channelId�UTF8c                 C   s
   || _ d S �N�r   ��selfr   � r   �2/storage/sdcard0/bot1/flex/curve/ChannelService.py�__init__   s    zissueChannelToken_args.__init__c                 C   s�   t jd k	r.| jd k	r.|�| || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tj	kr�t
jd dkr~|�� �d�n|�� | _q�|�|� n
|�|� |��  q6|��  d S �Nr   r   �   �utf-8)�oprot�_fast_encode�thrift_spec�_fast_decode�	__class__�readStructBegin�readFieldBeginr   �STOP�STRING�sys�version_info�
readString�decoder   �skip�readFieldEnd�readStructEnd�r   �iprot�fname�ftype�fidr   r   r   �read   s    

(

zissueChannelToken_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )Nr   r   r   r   r   r   �r   r   �trans�writer!   �writeStructBeginr   �writeFieldBeginr   r%   �writeStringr&   r'   �encode�writeFieldEnd�writeFieldStop�writeStructEnd�r   r   r   r   r   r5   ,   s    

&zissueChannelToken_args.write)N�
�__name__�
__module__�__qualname__�__doc__r   r%   r   r   r2   r5   r   r   r   r   r      s   �
r   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�issueChannelToken_result�-
    Attributes:
     - success
     - e
    r   �successNr   �ec                 C   s   || _ || _d S r   �rE   rF   �r   rE   rF   r   r   r   r   C   s    z!issueChannelToken_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S �Nr   r   �r    r   r!   r"   r#   r   r$   �STRUCTr   rE   r2   r*   r   rF   r+   r,   r-   r   r   r   r2   G   s(    




zissueChannelToken_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrC   rE   r   rF   r   �r   r   r4   r5   r!   r6   rE   r7   r   rK   r:   rF   r;   r<   r=   r   r   r   r5   a   s    


zissueChannelToken_result.write)NN�r?   r@   rA   rB   r   rK   r   r   r   r   r2   r5   r   r   r   r   rC   8   s   �
rC   c                   @   s>   e Zd ZdZddejdddffZddd�Zdd	� Zd
d� Z	dS )�'approveChannelAndIssueChannelToken_argsr   Nr   r   r   c                 C   s
   || _ d S r   r   r   r   r   r   r   |   s    z0approveChannelAndIssueChannelToken_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� n
|�|� |��  q6|��  d S r   �r    r   r!   r"   r#   r   r$   r%   r&   r'   r(   r)   r   r*   r+   r,   r-   r   r   r   r2      s    

(

z,approveChannelAndIssueChannelToken_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )NrN   r   r   r   r   r   r3   r=   r   r   r   r5   �   s    

&z-approveChannelAndIssueChannelToken_args.write)Nr>   r   r   r   r   rN   q   s   �
rN   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�)approveChannelAndIssueChannelToken_resultrD   r   rE   Nr   rF   c                 C   s   || _ || _d S r   rG   rH   r   r   r   r   �   s    z2approveChannelAndIssueChannelToken_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S rI   rJ   r-   r   r   r   r2   �   s(    




z.approveChannelAndIssueChannelToken_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrP   rE   r   rF   r   rL   r=   r   r   r   r5   �   s    


z/approveChannelAndIssueChannelToken_result.write)NNrM   r   r   r   r   rP   �   s   �
rP   c                   @   sN   e Zd ZdZdddejdddfdejdddffZddd	�Zd
d� Zdd� Z	dS )�getChannelInfo_argsz4
    Attributes:
     - channelId
     - locale
    Nr   r   r   �   �localec                 C   s   || _ || _d S r   )r   rS   )r   r   rS   r   r   r   r   �   s    zgetChannelInfo_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� nP|dkr�|tjkr�tj	d dkr�|�
� �d�n|�
� | _q�|�|� n
|�|� |��  q6|��  d S )Nr   r   r   rR   )r    r   r!   r"   r#   r   r$   r%   r&   r'   r(   r)   r   r*   rS   r+   r,   r-   r   r   r   r2   �   s$    

(
(

zgetChannelInfo_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  | jd k	r�|�dtj	d� |�
tjd dkr�| j�d�n| j� |��  |��  |��  d S )NrQ   r   r   r   r   rS   rR   )r   r   r4   r5   r!   r6   r   r7   r   r%   r8   r&   r'   r9   r:   rS   r;   r<   r=   r   r   r   r5     s    

&
&zgetChannelInfo_args.write)NNr>   r   r   r   r   rQ   �   s   �
rQ   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�getChannelInfo_resultrD   r   rE   Nr   rF   c                 C   s   || _ || _d S r   rG   rH   r   r   r   r     s    zgetChannelInfo_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S rI   )r    r   r!   r"   r#   r   r$   rK   r	   rE   r2   r*   r   rF   r+   r,   r-   r   r   r   r2   !  s(    




zgetChannelInfo_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrT   rE   r   rF   r   rL   r=   r   r   r   r5   ;  s    


zgetChannelInfo_result.write)NN)r?   r@   rA   rB   r   rK   r	   r   r   r   r2   r5   r   r   r   r   rT     s   �
rT   c                   @   sN   e Zd ZdZdddejdddfdejdddffZddd	�Zd
d� Z	dd� Z
dS )�getChannels_argsz5
    Attributes:
     - lastSynced
     - locale
    Nr   �
lastSyncedrR   rS   r   c                 C   s   || _ || _d S r   )rV   rS   )r   rV   rS   r   r   r   r   X  s    zgetChannels_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkrz|tjkrn|�� | _	q�|�
|� nP|dkr�|tjkr�tjd dkr�|�� �d�n|�� | _q�|�
|� n
|�
|� |��  q6|��  d S )Nr   rR   r   r   )r    r   r!   r"   r#   r   r$   �I64�readI64rV   r*   r%   r&   r'   r(   r)   rS   r+   r,   r-   r   r   r   r2   \  s$    


(

zgetChannels_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� |�
| j� |��  | jd k	r�|�dtjd� |�tjd dkr�| j�d�n| j� |��  |��  |��  d S )NrU   rV   r   rS   rR   r   r   )r   r   r4   r5   r!   r6   rV   r7   r   rW   �writeI64r:   rS   r%   r8   r&   r'   r9   r;   r<   r=   r   r   r   r5   t  s    


&zgetChannels_args.write)NN)r?   r@   rA   rB   r   rW   r%   r   r   r2   r5   r   r   r   r   rU   K  s   �
rU   c                   @   sR   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� ZdS )�getChannels_resultrD   r   rE   Nr   rF   c                 C   s   || _ || _d S r   rG   rH   r   r   r   r   �  s    zgetChannels_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  d S rI   )r    r   r!   r"   r#   r   r$   rK   r
   rE   r2   r*   r   rF   r+   r,   r-   r   r   r   r2   �  s(    




zgetChannels_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )NrZ   rE   r   rF   r   rL   r=   r   r   r   r5   �  s    


zgetChannels_result.write)NN)r?   r@   rA   rB   r   rK   r
   r   r   r   r2   r5   r   r   r   r   rZ   �  s   �
rZ   c                   @   s>   e Zd ZdZddejdddffZddd�Zdd	� Zd
d� Z	dS )�revokeChannel_argsr   Nr   r   r   c                 C   s
   || _ d S r   r   r   r   r   r   r   �  s    zrevokeChannel_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkr�tj	d dkr~|�
� �d�n|�
� | _q�|�|� n
|�|� |��  q6|��  d S r   rO   r-   r   r   r   r2   �  s    

(

zrevokeChannel_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	r�|�dtj	d� |�
tjd dkrv| j�d�n| j� |��  |��  |��  d S )Nr[   r   r   r   r   r   r3   r=   r   r   r   r5   �  s    

&zrevokeChannel_args.write)Nr>   r   r   r   r   r[   �  s   �
r[   c                   @   sB   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dS )�revokeChannel_resultz
    Attributes:
     - e
    Nr   rF   c                 C   s
   || _ d S r   )rF   )r   rF   r   r   r   r   �  s    zrevokeChannel_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S )Nr   )r    r   r!   r"   r#   r   r$   rK   r   rF   r2   r*   r+   r,   r-   r   r   r   r2   �  s    



zrevokeChannel_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr\   rF   r   )r   r   r4   r5   r!   r6   rF   r7   r   rK   r:   r;   r<   r=   r   r   r   r5   
  s    

zrevokeChannel_result.write)N)r?   r@   rA   rB   r   rK   r   r   r   r2   r5   r   r   r   r   r\   �  s   �
r\   c                   @   s>   e Zd Zddd�Zdd� Zdd� Zdd	� Zd
d� Zdd� ZdS )�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   r.   r   r   r   r   r     s    zClient.__init__c                 C   s�   | j �dtj| j� t|d��| j � | j ��  | j j�	�  | j
}|�� \}}}|tjkrvt� }|�|� |��  |�t� }|�|� |��  |jd k	r�|jS |jd k	r�t|j��ttjd��d S )N�issueChannelTokenr   z(issueChannelToken failed: unknown result)r_   �writeMessageBeginr   �CALLr`   r   r5   �writeMessageEndr4   �flushr^   �readMessageBegin�	EXCEPTIONr   r2   �readMessageEndrC   rE   rF   �	Exception�MISSING_RESULT�r   r   r.   r/   �mtype�rseqid�x�resultr   r   r   ra     s&    






zClient.issueChannelTokenc                 C   s�   | j �dtj| j� t|d��| j � | j ��  | j j�	�  | j
}|�� \}}}|tjkrvt� }|�|� |��  |�t� }|�|� |��  |jd k	r�|jS |jd k	r�t|j��ttjd��d S )N�"approveChannelAndIssueChannelTokenr   z9approveChannelAndIssueChannelToken failed: unknown result)r_   rb   r   rc   r`   rN   r5   rd   r4   re   r^   rf   rg   r   r2   rh   rP   rE   rF   ri   rj   rk   r   r   r   rp   2  s&    






z)Client.approveChannelAndIssueChannelTokenc           
      C   s�   | j �dtj| j� t� }||_||_|�| j � | j �	�  | j j
��  | j}|�� \}}}|tjkr�t� }|�|� |��  |�t� }	|	�|� |��  |	jd k	r�|	jS |	jd k	r�t|	j��ttjd��d S )N�getChannelInfoz%getChannelInfo failed: unknown result)r_   rb   r   rc   r`   rQ   r   rS   r5   rd   r4   re   r^   rf   rg   r   r2   rh   rT   rE   rF   ri   rj   )
r   r   rS   �argsr.   r/   rl   rm   rn   ro   r   r   r   rq   G  s,    






zClient.getChannelInfoc           
      C   s�   | j �dtj| j� t� }||_||_|�| j � | j �	�  | j j
��  | j}|�� \}}}|tjkr�t� }|�|� |��  |�t� }	|	�|� |��  |	jd k	r�|	jS |	jd k	r�t|	j��ttjd��d S )N�getChannelsz"getChannels failed: unknown result)r_   rb   r   rc   r`   rU   rV   rS   r5   rd   r4   re   r^   rf   rg   r   r2   rh   rZ   rE   rF   ri   rj   )
r   rV   rS   rr   r.   r/   rl   rm   rn   ro   r   r   r   rs   _  s,    






zClient.getChannelsc           	      C   s�   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  | j}|�� \}}}|tjkr|t� }|�|� |��  |�t� }|�|� |��  |jd k	r�t|j��d S )N�revokeChannel)r_   rb   r   rc   r`   r[   r   r5   rd   r4   re   r^   rf   rg   r   r2   rh   r\   rF   ri   )	r   r   rr   r.   r/   rl   rm   rn   ro   r   r   r   rt   w  s&    





zClient.revokeChannel)N)	r?   r@   rA   r   ra   rp   rq   rs   rt   r   r   r   r   r]     s   
r]   )�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   r&   �logging�ttypesr	   r
   r   r   r   �thrift.transportr   �objectr   rC   rN   rP   rQ   rT   rU   rZ   r[   r\   r]   r   r   r   r   �<module>   s"   ,9-::999,-